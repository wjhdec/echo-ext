# 基于 echo 的扩展

## 说明

对echo(v4)进行了功能的整合扩展，主要特点：


* 使用泛型封装了请求及返回，可以更方便的写 swag 注释，示例见 [main.go](./examples/main.go)
* 添加 form-data 类型的 binder，支持文件上传的绑定，示例见 [file_binder_test.go](./file_binder_test.go)
* 仿照 springboot 的默认错误结构重写 ErrorHandler


## 安装

```bash
go get -u github.com/wjhdec/echo-ext/v2
```

## 使用示例

见 [main.go](./examples/main.go)

测试 

`http://localhost:8181/my-test/sum?v1=1&v2=10`

测试错误：

`http://localhost:8181/my-test/demo-error`
