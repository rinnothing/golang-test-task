package server

import (
	"context"

	"go.uber.org/zap"

	"github.com/rinnothing/golang-test-task/api/gen"
	"github.com/rinnothing/golang-test-task/internal/usecase/integer"
	"github.com/rinnothing/golang-test-task/pkg/logger"
)

var _ gen.ServerInterface = &serverImplementation{}

type serverImplementation struct {
	l *zap.Logger

	integers integer.Usecase
}

func New(integers integer.Usecase, l *zap.Logger) *serverImplementation {
	return &serverImplementation{l: l, integers: integers}
}

func (s *serverImplementation) withLogger(ctx context.Context) context.Context {
	return logger.NewContext(ctx, s.l)
}
