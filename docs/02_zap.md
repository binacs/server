# zap & lumberjack

[GoDoc](https://godoc.org/go.uber.org/zap)



## 1. 概述

高性能日志库。



## 2. 关键概念

简要阅读参考：https://juejin.im/entry/5b741787e51d4566530c47a9



## 3. 使用示例

```go
func initLogger(rootPath string, logConfig config.LogConfig) *zap.Logger {
	logpath := logConfig.File
	if !path.IsAbs(logpath) {
		logpath = path.Join(rootPath, logConfig.File)
	}
  // lumberjack
  hook := lumberjack.Logger{...}
  // 编码器配置
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
  // core 传入: encoderConfig lumberjack 和 atomicLevel 等等
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(&hook), atomicLevel)
  // logger
	logger := zap.New(core)
	return logger
}
```

注意：

```go
type WriteSyncer interface {
    io.Writer
    Sync() error
}

func AddSync(w io.Writer) WriteSyncer
// AddSync converts an io.Writer to a WriteSyncer. It attempts to be intelligent: if the concrete type of the io.Writer implements WriteSyncer, we'll use the existing Sync method. If it doesn't, we'll add a no-op Sync.
```

这是实现日志分割的关键。



**在项目中对 zap Logger 进行了一定程度的封装，以支持 grpc、mysql 所需 Logger 以及 NopLogger 等实现。详情阅读 libs/log 模块代码。**



## 4. lumberjack

Lumberjack 用于将日志写入滚动文件。zap 不支持文件归档，如果要支持文件按大小或者时间归档，需要使用lumberjack，lumberjack 也是 zap 官方推荐的。

https://github.com/natefinch/lumberjack

如需同时打印到控制台和文件：

```go

zapcore.NewCore(
  	// Encoder 可以更改
		zapcore.NewJSONEncoder(encoderConfig),
  	// 请注意这里：
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
)
```

