package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"
)

// var (
// 	wg        sync.WaitGroup
// 	server1IP = "127.0.0.1:8088"
// 	server2IP = "127.0.0.1:8089"
// )

// func main() {
// 	r, err := nacos.NewDefaultNacosRegistry()
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		h := server.Default(
// 			server.WithHostPorts(server1IP),
// 			server.WithRegistry(r, &registry.Info{
// 				ServiceName: "hertz.test.demo",
// 				Addr:        utils.NewNetAddr("tcp", server1IP),
// 				Weight:      10,
// 				Tags:        nil,
// 			}),
// 		)
// 		h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
// 			ctx.JSON(consts.StatusOK, utils.H{"ping1": "pong1"})
// 		})
// 		h.Spin()
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		h := server.Default(
// 			server.WithHostPorts(server2IP),
// 			server.WithRegistry(r, &registry.Info{
// 				ServiceName: "hertz.test.demo",
// 				Addr:        utils.NewNetAddr("tcp", server2IP),
// 				Weight:      10,
// 				Tags:        nil,
// 			}),
// 		)
// 		h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
// 			ctx.JSON(consts.StatusOK, utils.H{"ping2": "pong2"})
// 		})
// 		h.Spin()
// 	}()

//		wg.Wait()
//	}
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		num := i
		go func() {
			addr := fmt.Sprintf("127.0.0.1:800%d", num)
			r, err := nacos.NewDefaultNacosRegistry()
			if err != nil {
				hlog.Fatal(err)
			}
			h := server.Default(
				server.WithHostPorts(addr),
				server.WithRegistry(r, &registry.Info{
					ServiceName: "hertz.test.demo",
					Addr:        utils.NewNetAddr("tcp", addr),
					Weight:      10,
					Tags:        nil,
				}),
			)
			h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
				ctx.JSON(consts.StatusOK, utils.H{"addr": addr})
			})
			h.Spin()
			wg.Done()
		}()
	}
	wg.Wait()
}
