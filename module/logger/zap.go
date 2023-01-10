package logger

import (
	"context"

	"budgetapp/module/consts"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	z *zap.SugaredLogger
}

func newZapLogger(env consts.Environment, opts Options) (Logger, error) {

	if env == consts.LOCAL {
		return getLocalEnvLogger()
	}

	return newLogger(opts)
}

func getLocalEnvLogger() (Logger, error) {
	option := zap.AddCallerSkip(1)

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := config.Build(option)
	if err != nil {
		return nil, err
	}
	return &zapLogger{
		z: l.Sugar(),
	}, nil
}

func newLogger(opts Options) (Logger, error) {
	loggerCfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	plain, err := loggerCfg.Build(zap.AddStacktrace(zapcore.Level(opts.MinStackTraceLvl)), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return &zapLogger{
		z: plain.Sugar()}, nil
}

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    encodeLevel(),
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}

func (l *zapLogger) ErrorW(ctx context.Context, message string, args ...interface{}) {
	args = appendRequestIDIntoArgs(ctx, args)
	l.z.Errorw(message, args...)
}

func (l *zapLogger) Errorf(ctx context.Context, message string, args ...interface{}) {
	l.z.Errorf(message, args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.z.Error(args...)
}

func (l *zapLogger) InfoW(ctx context.Context, message string, args ...interface{}) {
	args = appendRequestIDIntoArgs(ctx, args)
	l.z.Infow(message, args...)
}

func (l *zapLogger) Infof(ctx context.Context, message string, args ...interface{}) {
	l.z.Infof(message, args...)
}

func (l *zapLogger) Info(ctx context.Context, args ...interface{}) {
	l.z.Info(args...)
}

func (l *zapLogger) WarnW(ctx context.Context, message string, args ...interface{}) {
	args = appendRequestIDIntoArgs(ctx, args)
	l.z.Warnw(message, args...)
}

func (l *zapLogger) Warnf(ctx context.Context, message string, args ...interface{}) {
	l.z.Warnf(message, args...)
}

func (l *zapLogger) Warn(ctx context.Context, args ...interface{}) {
	l.z.Warn(args...)
}

func appendRequestIDIntoArgs(ctx context.Context, args []interface{}) []interface{} {
	ridValue, ok := ctx.Value(consts.RequestIDKey).(string)
	if !ok {
		return args
	}
	args = append(args, consts.RequestIDKey)
	args = append(args, ridValue)
	return args
}
