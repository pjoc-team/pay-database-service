package logger

import "github.com/pjoc-team/base-service/pkg/logger"

type LogrusLogger struct {
}

func (LogrusLogger) Print(v ...interface{}) {
	logger.Log.Info(v)
}
