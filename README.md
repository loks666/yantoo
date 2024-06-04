好的，以下是简单的步骤来在 Go 项目中使用 Swagger：

## 步骤 1：在接口中添加 Swagger 注释

在你的控制器文件中添加 Swagger 注释。例如，在 `controller/hello.go` 文件中：

```go
// Hello 控制器函数
// @Summary Hello endpoint
// @Description Returns a greeting message
// @ID hello
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response{data=string}
// @Router /api/hello [get]
func Hello(c *gin.Context) {
    c.JSON(200, common.Success("你好，登录成功"))
}
```

## 步骤 2：运行 `swag init` 命令生成 Swagger 文档

在项目根目录下运行以下命令生成 Swagger 文档：

```bash
swag init
```

## 步骤 3：重启项目并访问 Swagger UI

启动你的 Gin 服务器：

```bash
go run main.go
```

然后访问 `http://localhost:3000/doc/index.html`，你应该能够看到生成的 Swagger 文档。

通过这些简单的步骤，你可以在 Go 项目中使用 Swagger 生成和查看 API 文档。