# 基于 echo 的扩展

## 说明

对echo进行了功能的整合扩展，方便简单服务的开发

使用框架：

* echo
* zap
* sqlx
* viper
* sqldb-logger


## 使用示例

```go

package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wjhdec/echo-ext/pkg/elog"
	"github.com/wjhdec/echo-ext/pkg/server"
	"net/http"
)

type ResultInfo struct {
	Value string
}

type Req struct {
	Name  string  `query:"name"`
	Value float64 `query:"value"`
}

func NewTest1Handler() server.HandlerEnable {
	return server.NewJsonHandler("", http.MethodGet, func(req *Req) (*ResultInfo, error) {
		return &ResultInfo{Value: req.Name + "_" + fmt.Sprintf("%f", req.Value)}, nil
	})
}

func NewDemoRouter(group *echo.Group) *server.Router {
	router := server.NewRouter(group)
	router.AddHandler(NewTest1Handler())
	return router
}

func main() {
	svr, err := server.NewServer("v0.0")
	if err != nil {
		elog.Error(err)
	}
	svr.AddMiddleware(middleware.Logger())
	svr.AddRouter(NewDemoRouter(svr.RootGroup()))
	svr.Run()
}

```
