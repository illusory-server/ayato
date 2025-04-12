package logger

import (
	"context"
)

const (
	_ = iota
	StringType
	IntType
	AnyType
	ErrorType
	DurationType
	TimeType
	BoolType
	Float32Type
	Float64Type
	Int8Type
	Int16Type
	Int32Type
	Int64Type
	Uint8Type
	Uint16Type
	Uint32Type
	Uint64Type
	RawJsonType
	GroupType
)

type Level int

const (
	DebugLvl Level = -10
	InfoLvl  Level = 0
	WarnLvl  Level = 10
	ErrorLvl Level = 20
)

const (
	ErrKey = "error"
)

type (
	Logger interface {
		Log(ctx context.Context, level Level, message string, fields ...Field)
		Debug(ctx context.Context, message string, fields ...Field)
		Info(ctx context.Context, message string, fields ...Field)
		Warn(ctx context.Context, message string, fields ...Field)
		Error(ctx context.Context, message string, fields ...Field)
		With(fields ...Field) Logger
		InjectCtx(ctx context.Context) context.Context
		Enabled(ctx context.Context, level Level) bool
	}

	Field struct {
		Key   string
		Type  int
		Value any
	}

	NoopLogger struct{}
)

func (n NoopLogger) Log(context.Context, Level, string, ...Field) {}

func (n NoopLogger) Debug(context.Context, string, ...Field) {}

func (n NoopLogger) Info(context.Context, string, ...Field) {}

func (n NoopLogger) Warn(context.Context, string, ...Field) {}

func (n NoopLogger) Error(context.Context, string, ...Field) {}

func (n NoopLogger) With(...Field) Logger { return n }

func (n NoopLogger) InjectCtx(ctx context.Context) context.Context { return ctx }

func (n NoopLogger) Enabled(context.Context, Level) bool { return false }
