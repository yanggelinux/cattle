package log

import (
	commonLog "github.com/yanggelinux/cattle/common/log"
	"github.com/yanggelinux/cattle/global"
	"path"
	"strings"
)

var Logger commonLog.Logger

func SetupLog() {
	opts := &commonLog.Options{
		Level:             strings.ToLower(global.APPSetting.LogLevel),
		DisableCaller:     false,
		DisableStacktrace: true,
		Format:            commonLog.JsonFormat,
		EnableColor:       false,
		Development:       global.APPSetting.Development,
		OutputPath:        path.Join(global.APPSetting.LogPath, global.APPSetting.Service),
		Name:              global.APPSetting.Name,
		SName:             global.APPSetting.Service,
		HookOptions: &commonLog.HookOptions{
			MaxSize:    300,   // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 15,    // 日志文件最多保存多少个备份
			MaxAge:     15,    // 文件最多保存多少天
			Compress:   false, // 是否压缩
		},
	}
	Logger = commonLog.New(opts)

}
