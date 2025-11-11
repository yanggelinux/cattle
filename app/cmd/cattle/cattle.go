package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/yanggelinux/cattle/cron"
	"github.com/yanggelinux/cattle/global"
	"github.com/yanggelinux/cattle/internal/router"
	"github.com/yanggelinux/cattle/pkg/log"
	"github.com/yanggelinux/cattle/setup"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		config string
		err    error
	)
	flag.StringVar(&config, "config", "", "请制定配置文件路径")
	flag.Parse()

	if len(config) == 0 {
		config = "./configs/config.yaml"
	}
	//初始化配置
	err = setup.SetupAny(config)
	if err != nil {
		panic(err)
	}
	runMsg := fmt.Sprintf("config is %s", config)
	log.Logger.Info(runMsg)
	// 定时任务启动开关
	if global.CronSetting.Status == 1 {
		err = cron.InitCronTask()
		if err != nil {
			log.Logger.Error("init cron task error:", zap.Error(err))
			return
		}
	}

	//gin
	routerHandler := router.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.ServerSetting.HTTPPort),
		Handler:        routerHandler,
		ReadTimeout:    time.Duration(global.ServerSetting.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(global.ServerSetting.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	runMsg = fmt.Sprintf("running app success port %d,config is %s", global.ServerSetting.HTTPPort, config)
	log.Logger.Info(runMsg)

	//优雅起停服务
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		err := s.ListenAndServe()
		return err
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Logger.Info("errgroup exit...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Logger.Info("shutting down server...")
		return s.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("get os signal: %v", sig)
		}
	})
	log.Logger.Info("server start...")
	err = g.Wait()
	if err != nil {
		log.Logger.Error("http server exit...:", zap.Error(err))
	}
}
