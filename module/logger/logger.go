package logger

import (
	"context"

	"budgetapp/module/consts"
)

type Logger interface {
	InfoW(ctx context.Context, message string, args ...interface{})
	Infof(ctx context.Context, message string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, message string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	WarnW(ctx context.Context, message string, args ...interface{})
	Errorf(ctx context.Context, message string, args ...interface{})
	Error(args ...interface{})
	ErrorW(ctx context.Context, message string, args ...interface{})
}

type Factory interface {
	NewLogger() Logger
	NewLoggerWithOpts(opts Options) Logger
}

type factory struct {
	env     consts.Environment
	options Options
}

type Level int8

const (
	DebugLevel Level = iota - 1

	InfoLevel

	WarnLevel

	ErrorLevel

	DPanicLevel

	PanicLevel

	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	InvalidLevel = _maxLevel + 1
)

type Options struct {
	MinStackTraceLvl Level
}

var defaultOptions = Options{
	MinStackTraceLvl: ErrorLevel,
}

func NewFactory(env consts.Environment) Factory {
	return NewFactoryWithOpts(env, defaultOptions)
}

func NewFactoryWithOpts(env consts.Environment, opts Options) Factory {
	return &factory{
		env:     env,
		options: opts,
	}
}

func (f *factory) NewLogger() Logger {
	l, err := newZapLogger(f.env, f.options)

	if err != nil {
		panic("failed to initialize logger")
	}

	return l
}

func (f *factory) NewLoggerWithOpts(opts Options) Logger {
	l, err := newZapLogger(f.env, opts)

	if err != nil {
		panic("failed to initialize logger")
	}

	return l
}
