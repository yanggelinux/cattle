package setup

import (
	"github.com/yanggelinux/cattle/common/errors"
	"github.com/yanggelinux/cattle/global"
	"github.com/yanggelinux/cattle/internal/store/model"
	"github.com/yanggelinux/cattle/pkg/log"
	"github.com/yanggelinux/cattle/setting"
)

func SetupAny(config string) error {

	err := setupSetting(config)
	if err != nil {
		return err
	}
	//初始化Log配置
	log.SetupLog()
	err = model.SetupModel()
	if err != nil {
		return err
	}
	return nil
}

func setupSetting(config string) error {
	if config == "" {
		config = "configs/config.yaml"
	}
	sett, err := setting.NewSetting(config)
	if err != nil {
		err = errors.Wrap(err, "setting set new configs error")
		return err
	}
	err = sett.ReadSection("APP", &global.APPSetting)
	if err != nil {
		err = errors.Wrap(err, "setting read app section error")
		return err
	}
	err = sett.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		err = errors.Wrap(err, "setting server app section error")
		return err
	}
	err = sett.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		err = errors.Wrap(err, "setting read jwt section error")
		return err
	}
	err = sett.ReadSection("MySQL", &global.MySQLSetting)
	if err != nil {
		err = errors.Wrap(err, "setting read mysql section error")
		return err
	}

	err = sett.ReadSection("Cron", &global.CronSetting)
	if err != nil {
		err = errors.Wrap(err, "setting read cron section error")
		return err
	}

	return nil
}
