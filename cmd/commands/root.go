package commands

import (
	"path"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/binacsgo/log"

	"github.com/binacs/server/config"
)

var (
	configFile string
	zaplogger  *zap.Logger
	logger     log.Logger
	cfg        *config.Config
)

func init() {
	rootCmdFlags(RootCmd)
}

func rootCmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&configFile, "configFile", "config.toml", "config file (default is ./config.toml)")
}

var (
	// RootCmd the root command
	RootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root Command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			if cmd.Name() != StartCmd.Name() {
				return nil
			}

			if cfg, err = config.LoadFromFile(configFile); err != nil {
				return err
			}

			initLogger()
			logger.Info("Init finished")

			return nil
		},
	}
)

func initLogger() {
	logpath := cfg.LogConfig.File
	if !path.IsAbs(cfg.LogConfig.File) {
		logpath = path.Join(cfg.WorkSpace, cfg.LogConfig.File)
	}

	hook := lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    cfg.LogConfig.MaxSize,
		MaxBackups: cfg.LogConfig.MaxBackups,
		MaxAge:     cfg.LogConfig.MaxAge,
		Compress:   true, // Enable compression
		LocalTime:  true, // Use local time
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
	atomicLevel.SetLevel(stringToXZapLoggerLevel(cfg.LogConfig.Level))
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(&hook), atomicLevel)
	zaplogger = zap.New(core)

	logger = log.NewZapLoggerWrapper(zaplogger.Sugar())
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
	case "fatal":
		return zap.FatalLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}
