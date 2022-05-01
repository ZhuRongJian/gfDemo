package cmd

import (
	"context"
	"gfDemo/internal/service"
	"gfDemo/utility/captcha"
	"gfDemo/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gfDemo/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			//初始化验证码
			captcha.Setup()

			s := g.Server()
			s.Use(middleware.Middleware.CORS)

			//后端接口
			s.Group("/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Public,
					controller.Login,
				)
				group.Group("/", func(auth *ghttp.RouterGroup) {
					auth.Middleware(service.Token.Middleware)
					auth.Bind(
						controller.Index,
						controller.Menu,
						controller.Role,
						controller.Config,
						controller.Admin,
					)
					//auth.Middleware(service.Permission.Middleware)
				})
			})

			//前端接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind()
			})
			s.Run()
			return nil
		},
	}
)
