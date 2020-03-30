package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
	"strings"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/libs/log"
)

var (
	configFile string
	logger     log.Logger
	cfg        *config.Config
)

func init() {
	rootCmdFlags(RootCmd)
}

func rootCmdFlags(cmd *cobra.Command) {
	RootCmd.PersistentFlags().StringVar(&configFile, "configFile", "config.toml", "config file (default is ./config.toml)")
}

var (
	RootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root Command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if cmd.Name() != StartCmd.Name() {
				return nil
			}

			cfg, err = config.LoadFromFile(configFile)
			if err != nil {
				fmt.Println("LoadFromFile err: ", err)
			}

			logger = initLogger(cfg.WorkSpace, cfg.LogConfig)
			logger.Info("initLogger finished")

			return nil
		},
	}
)

func initLogger(rootPath string, logConfig config.LogConfig) log.Logger {
	logpath := logConfig.File
	if !path.IsAbs(logpath) {
		logpath = path.Join(rootPath, logConfig.File)
	}
	fmt.Printf("Log path : %s\n", logpath)
	hook := lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    logConfig.Maxsize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.Maxage,
		Compress:   true,
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "_time",
		LevelKey:       "_level",
		NameKey:        "_logger",
		CallerKey:      "_caller",
		MessageKey:     "_message",
		StacktraceKey:  "_stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(stringToXZapLoggerLevel(logConfig.Level))
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(&hook), atomicLevel)
	logger := zap.New(core)
	return log.NewZapLoggerWapper(logger.Sugar())
}

func stringToXZapLoggerLevel(level string) zapcore.Level {
	lower := strings.ToLower(level)
	switch lower {
	case "info":
		return zap.InfoLevel
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
