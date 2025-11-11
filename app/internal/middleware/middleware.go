package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/internal/pkg/app"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/pkg/log"
	"github.com/yanggelinux/cattle/pkg/util"
	"net/url"
	"strconv"
	"time"
)

const XRequestID = "X-RequestID"
const XUserame = "X-Username"
const XUserID = "X-UserID"
const XAuthorization = "X-Authorization"
const XSuper = "X-Super"
const XDeptName = "X-DeptName"

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token  string
			userID string
			err    error
		)

		data := make(map[string]interface{})
		s, exist := c.GetQuery("token")
		if exist {
			token = s
		} else {
			token = c.GetHeader("X-Token")
		}
		if token == "" {
			err = errors.WithCodeError(ce.ErrorAuthTokenFailed.Code(), errors.New("token is empty"))
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = errors.WithCodeError(ce.ErrorAuthCheckTokenTimeout.Code(), err)
				default:
					err = errors.WithCodeError(ce.ErrorAuthCheckTokenFailed.Code(), err)
				}
				if err != nil {
					app.Response(c, data, err)
					c.Abort()
					return
				}
			}
			if claims != nil {
				_userID := claims.UserID
				userID = strconv.Itoa(int(_userID))
			}
		}
		if err != nil {
			app.Response(c, data, err)
			c.Abort()
			return
		}
		c.Request.Header.Set(XUserID, userID)
		c.Set(XUserID, userID)
		c.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			data := make(map[string]interface{})
			if p := recover(); p != nil {
				//painc时打印出堆栈日志
				errMsg := fmt.Sprintf("%+v", p)
				err := errors.WithCode(ce.Error.Code(), "panic recover", errors.New(errMsg))
				app.Response(c, data, err)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func ReqCostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()

		//请求处理
		c.Next()

		//处理后获取消耗时间
		costTime := time.Since(nowTime)
		_url := c.Request.URL.String()
		msg := fmt.Sprintf("the request URL %s cost %v", _url, costTime)
		log.Logger.Info(msg)
	}
}

func RequestMetadata() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		var (
			//userID   string
			userName string
			rid      string
		)
		rid = c.GetHeader(XRequestID)
		if len(rid) == 0 {
			rid = util.GenUUIDv4()
		}
		_userName := c.GetHeader(XUserame)
		_deptName := c.GetHeader(XDeptName)
		_super := c.GetHeader(XSuper)
		_authorization := c.GetHeader(XAuthorization)
		//_userID := c.GetHeader(XUserID)
		//if len(_userID) == 0 {
		//	userID = "0"
		//}
		userName, err := url.QueryUnescape(_userName)
		if err != nil {
			userName = _userName
		}
		if len(userName) == 0 {
			userName = "other client"
		}
		c.Request.Header.Set(XRequestID, rid)
		c.Request.Header.Set(XUserame, userName)
		c.Request.Header.Set(XDeptName, _deptName)
		c.Request.Header.Set(XAuthorization, _authorization)
		//c.Request.Header.Set(XUserID, userID)
		c.Set(XRequestID, rid)
		//c.Set(XUserID, userID)
		c.Set(XUserame, userName)
		c.Set(XDeptName, _deptName)
		c.Set(XAuthorization, _authorization)
		c.Set(XSuper, _super)
		c.Next()
	}
}
