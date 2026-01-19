package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/rinnothing/golang-test-task/api/gen"
	"github.com/rinnothing/golang-test-task/internal/model"
	"github.com/rinnothing/golang-test-task/pkg/logger"
)

func (s *serverImplementation) PostIntegerAdd(c echo.Context) error {
	var body gen.PostIntegerAddJSONRequestBody
	if err := c.Bind(&body); err != nil {
		return badRequest(fmt.Sprintf("Invalid Request data in PostIntegerAddCreate: %s", err.Error()))
	}

	ctx := logger.NewContext(c.Request().Context(), s.l.With(zap.String("method", "PostIntegerAdd"), zap.Int("integer", body)))

	integer := model.Integer(body)
	sortedIntegers, err := s.integers.AddInteger(ctx, integer)
	if err != nil {
		return reportInternalError(ctx, err)
	}

	var resp []int
	for _, val := range sortedIntegers {
		resp = append(resp, int(val))
	}
	return c.JSON(http.StatusCreated, resp)
}
