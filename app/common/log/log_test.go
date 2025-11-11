package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLog(t *testing.T) {
	defer Flush() // used for record logger printer

	logger := WithName("test")
	logger.Infow("Hello world!", "foo", "bar") // structed logger

	logger = WithValues("key", "value") // used for record context
	logger.Info("Hello world!")
	logger.Info("Hello world!")

	logger = NewLogger(ZapLogger())
	logger.Info("test", zap.String("name", "yangyang"))

	opts := &Options{
		Level:             zapcore.ErrorLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: true,
		Format:            JsonFormat,
		EnableColor:       false,
		Development:       true,
		OutputPath:        "/tmp",
		Name:              "test",
		SName:             "log",
		HookOptions: &HookOptions{
			MaxSize:    1,     // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 15,    // 日志文件最多保存多少个备份
			MaxAge:     15,    // 文件最多保存多少天
			Compress:   false, // 是否压缩
		},
	}
	logger = New(opts)
	for i := 0; i < 100000; i++ {
		logger.Info("test info", zap.String("name", "yangyang"))
		logger.Error("test error", zap.String("name", "yangyang"))
	}

}
