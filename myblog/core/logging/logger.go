package logging

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

type LoggerIFace interface {
	InfoF(inCtx context.Context,format string, values... interface{})
}


type LoggerContext struct{
	logger *zap.Logger
	commonFields []zap.Field
}

func NewLoggerContext(logger *zap.Logger)*LoggerContext{
	return &LoggerContext{logger: logger}
}

func (l *LoggerContext) InfoF(inCtx context.Context,format string, values... interface{}){
	msg := fmt.Sprint(format,values)
	l.GenCommonFieldFromCtx(inCtx)
	l.logger.Info(msg,l.commonFields...)
}

func (l *LoggerContext) GenCommonFieldFromCtx(ctx context.Context){
	// 提取trace.id
	if traceIdI,ok := ctx.Value("trace.id").(string);ok{
		l.commonFields = append(l.commonFields,zap.String("trace.id",traceIdI))
	}
}

