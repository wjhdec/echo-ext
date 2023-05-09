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

见 [main.go](./examples/demo-server/main.go)

测试 `http://localhost:8888/my-test?name=a&value=200`