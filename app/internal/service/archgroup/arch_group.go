package archgroup

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type ArchGroupService interface {
	GetList(context.Context, *request.GetArchGroupReq) (*result.ArchGroupResult, error)
	Create(context.Context, *request.CreateArchGroupReq) error
	Update(context.Context, *request.UpdateArchGroupReq) error
	Delete(context.Context, int64) error
}
