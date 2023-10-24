package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/wjhdec/echo-ext/pkg/customfmt/custime"
	"github.com/wjhdec/echo-ext/pkg/logext"
	"github.com/wjhdec/echo-ext/pkg/set"
)

// ErrResponse 返回的错误结构
type ErrResponse struct {
	Timestamp custime.FormatTime `json:"timestamp"`
	Status    int                `json:"status"`
	Error     string             `json:"error"`
	Message   string             `json:"message"`
	Path      string             `json:"path"`
}

func NewErrResponse(err *echo.HTTPError, c echo.Context) *ErrResponse {
	return &ErrResponse{
		Timestamp: custime.Now(),
		Status:    err.Code,
		Error:     http.StatusText(err.Code),
		Message:   fmt.Sprintf("%s", err.Message),
		Path:      c.Request().RequestURI,
	}
}

func CustomHttpErrorHandler(err error, c echo.Context) {
	if !c.Response().Committed {
		if err := c.JSON(getErrorResponse(err, c)); err != nil {
			logext.LogError(err)
		}
	}
}

// getErrorResponse 获取返回，返回中int为status
func getErrorResponse(err error, c echo.Context) (int, any) {
	ignoreCode := set.NewWithSlice([]int{401, 403, 404})
	switch e := err.(type) {
	case *echo.HTTPError:
		if !ignoreCode.Contains(e.Code) {
			logext.LogError(err)
		}
		return e.Code, NewErrResponse(e, c)
	default:
		logext.LogError(err)
		code := http.StatusInternalServerError
		he := &echo.HTTPError{
			Code: code, Message: err.Error(),
		}
		return he.Code, NewErrResponse(he, c)
	}
}
