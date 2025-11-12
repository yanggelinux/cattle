package archgraph

import (
	"context"
	"github.com/yanggelinux/cattle/internal/dto/request"
	"github.com/yanggelinux/cattle/internal/dto/result"
)

type ArchGraphService interface {
	GetList(context.Context, *request.GetArchGraphReq) (*result.ArchGraphResult, error)
	Create(context.Context, *request.CreateArchGraphReq) (*result.ArchGraphOptResult, error)
	Update(context.Context, *request.UpdateArchGraphReq) error
	Save(context.Context, *request.UpdateArchGraphReq) error
	Delete(context.Context, int64) error
	GetByID(context.Context, int64) (*result.ArchGraphDetailResult, error)
	Copy(context.Context, int64) error

	// 记录相关
	GetRecordList(context.Context, *request.GetArchGraphRecordReq) (*result.ArchGraphRecordResult, error)
	SelectRecord(context.Context, *request.SelectArchGraphReq) error
	GetEnabledRecord(context.Context, int64) (*result.ArchGraphRecordRet, error)

	// 评审相关
	GetReviewList(context.Context, *request.GetArchGraphReviewReq) (*result.ArchGraphReviewResult, error)
	CreateReview(context.Context, *request.CreateArchGraphReviewReq) error
	DeleteReview(context.Context, int64) error

	GetDataByLabel(context.Context, *request.GetArchGraphDataReq) (*result.ArchGraphData, error)
}
