package archgraph

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ArchGraphReviewRepo interface {
	WithTx(*model.MDB) ArchGraphReviewRepo
	GetList(context.Context, *ArchGraphReviewFilter) ([]*model.ArchGraphReview, int64, error)
	GetByID(context.Context, int64) (*model.ArchGraphReview, error)
	Create(context.Context, *model.ArchGraphReview) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.ArchGraphReview) error
	DeleteByID(context.Context, int64) error
}
