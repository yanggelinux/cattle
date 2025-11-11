package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yanggelinux/cattle/common/coder"
	"github.com/yanggelinux/cattle/internal/dto/result"
	"github.com/yanggelinux/cattle/internal/pkg/ce"
	"github.com/yanggelinux/cattle/pkg/log"
	"github.com/yanggelinux/cattle/pkg/util"
	"go.uber.org/zap"
	"net/http"
)

func Response(ctx *gin.Context, data interface{}, err error) {
	// interface{} 为 nil 只有当类型和值都为 nil 才是“真正的 nil
	if data == nil {
		data = &result.EmptyResult{}
	}
	if err != nil {
		c := coder.ParseCoder(err)
		log.Logger.Error("请求错误", zap.Int("code", c.Code()), zap.String("msg", c.String()),
			zap.String("traceID", util.GetRequestID(ctx)), zap.Error(err))
		ctx.JSON(c.HTTPStatus(), gin.H{
			"status": c.Code(),
			"data":   data,
			"msg":    c.String(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": ce.Success.Code(),
			"data":   data,
			"msg":    ce.Success.String(),
		})
	}
}

func ResponseWithMsg(ctx *gin.Context, data interface{}, err error, msg string) {
	if err != nil {
		c := coder.ParseCoder(err)
		log.Logger.Error("请求错误", zap.Int("code", c.Code()), zap.String("msg", c.String()),
			zap.String("traceID", util.GetRequestID(ctx)), zap.Error(err))
		if msg == "" {
			msg = c.String()
		}
		ctx.JSON(c.HTTPStatus(), gin.H{
			"status": c.Code(),
			"data":   data,
			"msg":    msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": ce.Success.Code(),
			"data":   data,
			"msg":    ce.Success.String(),
		})
	}
}

func ResponseWithError(ctx *gin.Context, data interface{}, err error) {
	if err != nil {
		c := coder.ParseCoder(err)
		log.Logger.Error("请求错误", zap.Int("code", c.Code()), zap.String("msg", c.String()),
			zap.String("traceID", util.GetRequestID(ctx)), zap.Error(err))
		ctx.JSON(c.HTTPStatus(), gin.H{
			"status": c.Code(),
			"data":   data,
			"msg":    c.String() + ":" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": ce.Success.Code(),
			"data":   data,
			"msg":    ce.Success.String(),
		})
	}
}
