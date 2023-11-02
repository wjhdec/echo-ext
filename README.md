# 基于 echo 的扩展

## 说明

对echo进行了功能的整合扩展，方便简单服务的开发

依赖：

* github.com/labstack/echo/v4

## 安装

```bash
go get -u github.com/wjhdec/echo-ext/v2
```

## 使用示例

见 [main.go](./examples/main.go)

测试 

`http://localhost:8181/my-test/sum?v1=1&v2=10`

测试错误：

`http://localhost:8888/my-test/demo-error`
