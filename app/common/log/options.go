package log

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

const (
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-paths"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	ConsoleFormat = "console"
	JsonFormat    = "json"
)

// Options contains configuration items related to log.
type Options struct {
	OutputPath        string       `json:"output-paths"       mapstructure:"output-paths"`
	ErrorOutputPath   string       `json:"error-output-paths" mapstructure:"error-output-paths"`
	Level             string       `json:"level"              mapstructure:"level"`
	Format            string       `json:"format"             mapstructure:"format"`
	DisableCaller     bool         `json:"disable-caller"     mapstructure:"disable-caller"`
	DisableStacktrace bool         `json:"disable-stacktrace" mapstructure:"disable-stacktrace"`
	EnableColor       bool         `json:"enable-color"       mapstructure:"enable-color"`
	Development       bool         `json:"development"        mapstructure:"development"`
	LogStdout         bool         `json:"Log_stdout"        mapstructure:"Log_stdout"`
	Name              string       `json:"name"               mapstructure:"name"`
	SName             string       `json:"sname"               mapstructure:"sname"`
	HookOptions       *HookOptions `json:"hook_options" mapstructure:"hook_options"`
}

type HookOptions struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// NewOptions creates an Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		Level:             zapcore.ErrorLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: false,
		Format:            JsonFormat,
		EnableColor:       false,
		Development:       false,
		OutputPath:        "/tmp/zap",
		HookOptions: &HookOptions{
			MaxSize:    300,   // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 15,    // 日志文件最多保存多少个备份
			MaxAge:     15,    // 文件最多保存多少天
			Compress:   false, // 是否压缩
		},
	}
}

// Validate validate the options fields.
func (o *Options) Validate() []error {
	var errs []error

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		errs = append(errs, err)
	}

	format := strings.ToLower(o.Format)
	if format != ConsoleFormat && format != JsonFormat {
		errs = append(errs, fmt.Errorf("not a valid log format: %q", o.Format))
	}

	return errs
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000"))
}

func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}
