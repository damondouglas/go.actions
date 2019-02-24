package appengine

import (
	"context"
	"net/http"

	"github.com/damondouglas/go.actions/v2/logger"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type loggerBase struct {
	ctx context.Context
}

// New Logger using appengine log.
func New(r *http.Request) logger.Logger {
	return &loggerBase{
		ctx: appengine.NewContext(r),
	}
}

// Infof log to console.
func (l *loggerBase) Infof(format string, args ...interface{}) {
	log.Infof(l.ctx, format, args...)
}

// Errorf log to console.
func (l *loggerBase) Errorf(format string, args ...interface{}) {
	log.Errorf(l.ctx, format, args...)
}
