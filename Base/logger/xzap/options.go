package xzap

import (
	"context"
	logging "github.com/ProjectsTask/Base/logger"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

type Options struct {
	shouldLog        logging.Decider
	codeFunc         logging.ErrorToCode
	levelFunc        CodeToLevel
	durationFunc     DurationToField
	messageFunc      MessageProducer
	timestampForamat string
}
type Option func(*Options)

type CodeToLevel func(code codes.Code) zapcore.Level

type DurationToField func(duration time.Duration) zapcore.Field

// DefaultCodeToLevel 根据RPC服务端返回码返回zap日志级别
func DefaultCodeToLevel(code codes.Code) zapcore.Level {
	switch code {
	case codes.OK:
		return zap.InfoLevel
	case codes.Canceled:
		return zap.InfoLevel
	case codes.Unknown:
		return zap.ErrorLevel
	case codes.InvalidArgument:
		return zap.InfoLevel
	case codes.DeadlineExceeded:
		return zap.WarnLevel
	case codes.NotFound:
		return zap.InfoLevel
	case codes.AlreadyExists:
		return zap.InfoLevel
	case codes.PermissionDenied:
		return zap.WarnLevel
	case codes.Unauthenticated:
		return zap.InfoLevel // unauthenticated requests can happen
	case codes.ResourceExhausted:
		return zap.WarnLevel
	case codes.FailedPrecondition:
		return zap.WarnLevel
	case codes.Aborted:
		return zap.WarnLevel
	case codes.OutOfRange:
		return zap.WarnLevel
	case codes.Unimplemented:
		return zap.ErrorLevel
	case codes.Internal:
		return zap.ErrorLevel
	case codes.Unavailable:
		return zap.WarnLevel
	case codes.DataLoss:
		return zap.ErrorLevel
	default:
		if code >= 7000 {
			return zap.InfoLevel
		}

		return zap.ErrorLevel
	}
}

// DefaultClientCodeToLevel 根据RPC客户端返回码返回zap日志级别
func DefaultClientCodeToLevel(code codes.Code) zapcore.Level {
	switch code {
	case codes.OK:
		return zap.DebugLevel
	case codes.Canceled:
		return zap.DebugLevel
	case codes.Unknown:
		return zap.InfoLevel
	case codes.InvalidArgument:
		return zap.DebugLevel
	case codes.DeadlineExceeded:
		return zap.InfoLevel
	case codes.NotFound:
		return zap.DebugLevel
	case codes.AlreadyExists:
		return zap.DebugLevel
	case codes.PermissionDenied:
		return zap.InfoLevel
	case codes.Unauthenticated:
		return zap.InfoLevel // unauthenticated requests can happen
	case codes.ResourceExhausted:
		return zap.DebugLevel
	case codes.FailedPrecondition:
		return zap.DebugLevel
	case codes.Aborted:
		return zap.DebugLevel
	case codes.OutOfRange:
		return zap.DebugLevel
	case codes.Unimplemented:
		return zap.WarnLevel
	case codes.Internal:
		return zap.WarnLevel
	case codes.Unavailable:
		return zap.WarnLevel
	case codes.DataLoss:
		return zap.WarnLevel
	default:
		if code >= 7000 {
			return zap.InfoLevel
		}

		return zap.InfoLevel
	}
}

type MessageProducer func(ctx context.Context, msg string, level zapcore.Level, err error, fields []zapcore.Field)
