package archgraph

import (
	"context"
	"github.com/yanggelinux/cattle/internal/store/model"
)

type ArchGraphRecordRepo interface {
	WithTx(*model.MDB) ArchGraphRecordRepo
	GetList(context.Context, *ArchGraphRecordFilter) ([]*model.ArchGraphRecord, error)
	GetSnapshotRecord(context.Context, int64) (*model.ArchGraphRecord, error)
	GetByID(context.Context, int64) (*model.ArchGraphRecord, error)
	GetSyncRecord(context.Context, int64) (*model.ArchGraphRecord, error)
	GetEnabledRecords(context.Context, int64) ([]*model.ArchGraphRecord, error)
	Create(context.Context, *model.ArchGraphRecord) error
	Update(context.Context, int64, map[string]interface{}) error
	CreateOrUpdate(context.Context, *model.ArchGraphRecord) error
	DeleteByID(context.Context, int64) error
}
