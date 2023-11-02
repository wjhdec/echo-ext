package echoext

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/labstack/echo/v4"
)

type ErrResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
}

func NewErrResponse(err *echo.HTTPError, c echo.Context) *ErrResponse {
	return &ErrResponse{
		Timestamp: time.Now(),
		Status:    err.Code,
		Error:     http.StatusText(err.Code),
		Message:   fmt.Sprintf("%s", err.Message),
		Path:      c.Request().RequestURI,
	}
}

func CustomHttpErrorHandler(err error, c echo.Context) {
	if !c.Response().Committed {
		if err := c.JSON(getErrorResponse(err, c)); err != nil {
			logError(err)
		}
	}
}

// getErrorResponse 获取返回，返回中int为status
func getErrorResponse(err error, c echo.Context) (int, any) {
	ignoreCodes := []int{401, 403, 404}
	switch e := err.(type) {
	case *echo.HTTPError:
		if !slices.Contains(ignoreCodes, e.Code) {
			logError(err)
		}
		return e.Code, NewErrResponse(e, c)
	default:
		logError(err)
		code := http.StatusInternalServerError
		he := &echo.HTTPError{
			Code: code, Message: err.Error(),
		}
		return he.Code, NewErrResponse(he, c)
	}
}
