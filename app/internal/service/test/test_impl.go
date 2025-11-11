package test

import (
	"context"
	rolerepo "github.com/yanggelinux/cattle/internal/repository/role"
	userrepo "github.com/yanggelinux/cattle/internal/repository/user"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type TestService struct {
	userRepo userrepo.UserRepo
	roleRepo rolerepo.RoleRepo
}

func NewTestService() ITestService {
	return &TestService{
		roleRepo: rolerepo.NewRoleRepo(),
		userRepo: userrepo.NewUserRepo(),
	}
}

func (s *TestService) DoTestStatus(ctx context.Context) (map[string]interface{}, error) {
	_, err := s.userRepo.GetByID(ctx, 1)
	if err != nil {
		return nil, err
	}
	_, err = s.roleRepo.GetByID(ctx, 1)
	if err != nil {
		return nil, err
	}
	// 事务情况,返回一个支持tx 的repo
	tx := model.GetDB().WithContext(ctx).Begin()
	defer model.RecoverRollback(tx)

	userTx := s.userRepo.WithTx(tx)
	err = userTx.DeleteByID(ctx, 1)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	roleTx := s.roleRepo.WithTx(tx)
	err = roleTx.DeleteByID(ctx, 111)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return nil, nil
}
