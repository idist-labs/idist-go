package loggerProvider

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"idist-core/app/providers/configProvider"
	"os"
	"path/filepath"
)

var Logger *zap.Logger

func Init() {
	fmt.Println("------------------------------------------------------------")
	var err error
	cf := configProvider.GetConfig()

	if dir, err := os.Getwd(); err != nil {
		panic(err)
	} else if fi, err := os.Stat(filepath.Join(dir, "logs")); os.IsNotExist(err) || !fi.IsDir() {
		_ = os.MkdirAll(filepath.Join(dir, "logs"), 0755)
	}

	if cf.GetString("log.mode") == "console" {
		cfg := zap.NewDevelopmentConfig()
		cfg.Encoding = "console"
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stderr"}
		switch cf.GetString("log.level") {
		case "debug":
			cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
			break
		case "info":
			cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
			break
		case "warning":
			cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
			break
		case "error":
			cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
			break
		case "dpanic":
			cfg.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
			break
		case "panic":
			cfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
			break
		}

		Logger, err = cfg.Build()
		if err != nil {
			panic(err)
		}
	} else {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "logs/log.yaml",
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     30, // days
		})
		encoder := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), w, zap.InfoLevel)
		Logger = zap.New(core, zap.Development())
	}

}
