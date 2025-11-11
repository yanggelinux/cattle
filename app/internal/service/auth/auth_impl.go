package auth

import (
	"context"
	"github.com/samber/lo"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/global"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	permrepo "github.com/yanggelinux/cattle/internal/repository/permission"
	rolerepo "github.com/yanggelinux/cattle/internal/repository/role"
	userrepo "github.com/yanggelinux/cattle/internal/repository/user"
	"github.com/yanggelinux/cattle/internal/service/user"
	"strings"
	"time"
)

type authService struct {
	userRepo       userrepo.UserRepo
	roleRepo       rolerepo.RoleRepo
	permissionRepo permrepo.PermissionRepo

	userService user.UserService
}

func NewAuthService() AuthService {
	return &authService{
		userRepo:       userrepo.NewUserRepo(),
		roleRepo:       rolerepo.NewRoleRepo(),
		permissionRepo: permrepo.NewPermissionRepo(),
		userService:    user.NewUserService(),
	}
}

func (s *authService) Login(ctx context.Context, req *request.LoginReq) (*result.LoginResult, error) {

	var userID int64
	resultData := &result.LoginResult{}
	userName := *req.UserName
	//获取一个认证加密的信息
	password := *req.Password
	// 现从系统用户查询，查询不到再ldap认证
	sysUser, err := s.userRepo.GetBySysName(ctx, userName)
	if err != nil {
		return resultData, errors.WithCodeError(ce.ErrorLoginUserPasswordEmpty.Code(), err)
	}
	if password != sysUser.Password {
		err := errors.WithCodeError(ce.ErrorLoginUserFailed.Code(), err)
		return nil, err
	}
	userID = sysUser.ID
	resultData.UserID = sysUser.ID
	resultData.UserName = userName
	resultData.DisplayName = sysUser.DisplayName
	resultData.Email = sysUser.Email
	resultData.DeptName = sysUser.DeptName
	// 更新最后登录时间
	updateFields := make(map[string]interface{})
	updateFields[sysUser.LastLoginTimeField()] = time.Now()
	err = s.userRepo.Update(ctx, sysUser.ID, updateFields)
	if err != nil {
		err := errors.WithCodeError(ce.ErrorLoginFailed.Code(), err)
		return nil, err
	}

	token, err := app.GenerateToken(global.JWTSetting.JwtSecret, global.JWTSetting.JwtIssuer, userID)
	if err != nil {
		err := errors.WithCodeError(ce.ErrorAuthToken.Code(), err)
		return nil, err
	}
	rolePermResult, err := s.GetRolesPerms(ctx, userID)
	if err != nil {
		err := errors.WithCodeError(ce.ErrorLoginFailed.Code(), err)
		return nil, err
	}

	resultData.Token = token
	resultData.RolePermResult = rolePermResult
	resultData.Project = "cattle"
	return resultData, nil
}

// 获取用户Token
func (s *authService) GetToken(ctx context.Context, req *request.GetTokenReq) (*result.TokenResult, error) {
	userName := *req.UserName
	record, err := s.userRepo.GetByName(ctx, userName)
	if err != nil {
		err := errors.WithCodeError(ce.ErrorGetToken.Code(), err)
		return nil, err
	}
	userID := record.ID
	token, err := app.GenerateToken(global.JWTSetting.JwtSecret, global.JWTSetting.JwtIssuer, userID)
	if err != nil {
		err := errors.WithCodeError(ce.ErrorAuthToken.Code(), err)
		return nil, err
	}
	resultData := &result.TokenResult{}
	resultData.Token = token
	return resultData, nil
}

// 获取权限的接口
func (s *authService) GetUserPermList(ctx context.Context, req *request.GetUserPermReq) (*result.RolePermResult, error) {
	userID := *req.UserID
	rolePermResult, err := s.GetRolesPerms(ctx, userID)
	if err != nil {
		err := errors.WithCodeError(ce.Error.Code(), err)
		return nil, err
	}
	return rolePermResult, nil
}

func (s *authService) GetRolesPerms(ctx context.Context, userID int64) (*result.RolePermResult, error) {
	// 验证成功后获取对应的角色和权限信息
	var issuper int8
	roleNames := make([]string, 0)
	roleDisplayNames := make([]string, 0)
	menus := make([]string, 0)
	uris := make([]string, 0)
	roleRecords, err := s.roleRepo.GetRoles(ctx, userID)
	if err != nil {
		err := errors.WithCodeError(ce.ErrorLoginFailed.Code(), err)
		return nil, err
	}
	for _, roleRecord := range roleRecords {
		// 如果是超级管理员拥有所有权限
		isSuper := roleRecord.IsSuper

		roleID := roleRecord.ID
		roleName := roleRecord.RoleName
		displayName := roleRecord.DisplayName
		roleNames = append(roleNames, roleName)
		roleDisplayNames = append(roleDisplayNames, displayName)
		// 一个用户多个角色的时候，如果有一个角色是超级管理员，这个用户就是超级管理员
		if isSuper == 1 {
			issuper = 1
			continue
		}
		permRecords, err := s.permissionRepo.GetPermsByRole(ctx, roleID)
		if err != nil {
			err = errors.WithCodeError(ce.ErrorLoginFailed.Code(), err)
			return nil, err
		}
		// 根据角色 获取权限相关
		for _, permRecord := range permRecords {
			permType := permRecord.PermType
			// 菜单权限
			if permType == 1 {
				code := permRecord.Code
				if !lo.Contains(menus, code) {
					menus = append(menus, code)
				}
				// api权限
			} else if permType == 2 {
				uri := permRecord.Uri
				method := permRecord.Method
				newUri := uri + ":" + strings.ToLower(method)
				if !lo.Contains(uris, newUri) {
					uris = append(uris, newUri)
				}
			} else {
				continue
			}
		}
	}
	resultData := &result.RolePermResult{}
	resultData.IsSuper = issuper
	resultData.RoleNames = strings.Join(roleNames, ",")
	resultData.RoleDisplayNames = strings.Join(roleDisplayNames, ",")
	resultData.Menus = menus
	resultData.Uris = uris
	return resultData, err
}
