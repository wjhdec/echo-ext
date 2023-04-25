package server

import (
	"echoext/pkg/customfmt/custime"
	"echoext/pkg/elog"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
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
		Timestamp: custime.FormatTime(time.Now()),
		Status:    err.Code,
		Error:     http.StatusText(err.Code),
		Message:   fmt.Sprintf("%s", err.Message),
		Path:      c.Request().RequestURI,
	}
}

// CustomHttpErrorHandler 自定义错误处理
func CustomHttpErrorHandler(err error, c echo.Context) {
	log := elog.Default()
	log.Errorf("%+v", err)
	if !c.Response().Committed {
		if err := c.JSON(getErrorResponse(err, c)); err != nil {
			log.Errorf("%+v", err)
		}
	} else {
		log.Warn("already committed")
	}
}

// getErrorResponse 获取返回，返回中int为status
func getErrorResponse(err error, c echo.Context) (int, interface{}) {
	switch e := err.(type) {
	case *echo.HTTPError:
		return e.Code, NewErrResponse(e, c)
	default:
		he := &echo.HTTPError{
			Code: http.StatusInternalServerError, Message: err.Error(),
		}
		return he.Code, NewErrResponse(he, c)
	}
}
