package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/yanggelinux/cattle/internal/service/demand"
	"github.com/yanggelinux/cattle/pkg/log"
	"go.uber.org/zap"
)

// 检查生成需求评价
func checkDemandEvaluation() {
	svc := demand.NewDemandService()
	err := svc.CheckEvaluation(context.Background())
	if err != nil {
		log.Logger.Error("cron:exec check evaluation error", zap.Error(err))
	}
}

func InitCronTask() error {
	var err error
	c := cron.New()
	_, err = c.AddFunc("0 * * * *", checkDemandEvaluation)
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
