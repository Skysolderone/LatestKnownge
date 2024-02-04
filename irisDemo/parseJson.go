package main

import (
	"github.com/kataras/iris/v12"
)

// 定义结构体，反映 JSON 数据的结构
type YourStruct struct {
	Symbols []string `json:"symbols"`
	KindId  uint     `json:"kindId"`
	Dir     uint8    `json:"dir"`
}

func main() {
	app := iris.New()

	// 定义路由处理函数
	app.Post("/parseJSON", func(ctx iris.Context) {
		// 从请求中读取 JSON 数据并解析到结构体中
		var data YourStruct
		if err := ctx.ReadJSON(&data); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString("无法解析 JSON 数据：" + err.Error())
			return
		}

		// 打印解析后的数据
		ctx.JSON(data)
	})

	// 启动 Iris 服务器
	app.Run(iris.Addr(":8080"))
}
