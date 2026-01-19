package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/rinnothing/golang-test-task/pkg/logger"
)

func badRequest(reason string) error {
	return echo.NewHTTPError(http.StatusBadRequest, reason)
}

func internalError() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func reportInternalError(ctx context.Context, err error) error {
	logger.ErrorCtx(ctx, "got internal error", zap.Error(err))
	return internalError()
}
