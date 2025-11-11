package process

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type ProcessService interface {
	GetList(context.Context, *request.GetProcessReq) (*result.ProcessResult, error)
	GetDetail(context.Context, int64) (*result.ProcessDetailRet, error)
	Create(context.Context, *request.CreateProcessReq) (*result.ProcessOptResult, error)
	Update(context.Context, *request.UpdateProcessReq) error
	Delete(context.Context, int64) error
	Copy(context.Context, int64) error
}
