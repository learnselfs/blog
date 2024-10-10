package cmd

import (
	"blog/internal/consts"
	"blog/internal/controller/config"
	"blog/internal/controller/login"
	"blog/internal/controller/menu"
	"blog/internal/controller/rbac"
	"blog/internal/controller/token"
	"blog/internal/controller/user"
	"blog/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.BindHandler("GET:/hello", func(req *ghttp.Request) {
				api := s.GetOpenApi()
				//openapi := api.OpenAPI
				//req.Response.Writeln("<h1>Hello World!</h1>")
				req.Response.WriteJson(api)
			})

			//s.BindHandler(path, func(req *ghttp.Request) {
			//	err := req.Response.WriteTpl(tpl)
			//	if err != nil {
			//		req.Response.Writeln(err)
			//	}
			//})
			s.BindHandler(consts.PathLogin, consts.HandleLogin)

			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(middleware.MiddlewareJwt)

				s.Group(consts.AdminPrefix, func(group *ghttp.RouterGroup) {
					group.GET(consts.PathException403, consts.HandleException403)
					group.GET(consts.PathException404, consts.HandleException404)
					group.GET(consts.PathException500, consts.HandleException500)
				})

				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(service.Middleware().Auth)
				group.Bind(
					//test.NewV1(),
					token.NewV1(),
				)
			})

			s.Group("/api", func(api *ghttp.RouterGroup) {
				api.Group("/v1", func(v1 *ghttp.RouterGroup) {
					v1.Middleware(ghttp.MiddlewareHandlerResponse)
					v1.Bind(
						login.NewV1(),
					)
				})

				api.Group("/v2", func(v2 *ghttp.RouterGroup) {
					//v2.Middleware(middleware.MiddlewareJwt)
					v2.Middleware(ghttp.MiddlewareHandlerResponse)
					v2.Bind(
						user.NewV1(),
						rbac.NewV1(),
						config.NewV1(),
						menu.NewV1(),
					)
				})

			})

			s.Group(consts.AdminPrefix, func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().Auth)
				//group.Middleware(middleware.MiddlewareJwt)

				group.GET(consts.PathIndex, consts.HandleIndex)
				group.GET(consts.PathAnalysisIndex, consts.HandleAnalysisIndex)
				group.GET(consts.PathComponentAdmin, consts.HandleComponentAdmin)
				group.GET(consts.PathComponentGrid, consts.HandleComponentGrid)
				group.GET(consts.PathComponentNprogress, consts.HandleComponentNprogress)
				group.GET(consts.PathComponentToast, consts.HandleComponentToast)
				group.GET(consts.PathConsoleIndex, consts.HandleConsoleIndex)
				group.GET(consts.PathListingTable, consts.HandleListingTable)
				group.GET(consts.PathProfileIndex, consts.HandleProfileIndex)
				group.GET(consts.PathResultError, consts.HandleResultError)
				group.GET(consts.PathResultSuccess, consts.HandleResultSuccess)
				group.GET(consts.PathRoleBaseAccessControl, consts.HandleRoleBaseAccessControl)
				group.GET(consts.PathSystemMenu, consts.HandleSystemMenu)
				group.GET(consts.PathSystemUser, consts.HandleSystemUser)
				group.GET(consts.PathSystemRole, consts.HandleSystemRole)
				group.GET(consts.PathSystemPermission, consts.HandleSystemPermission)
				group.GET(consts.PathUtilTable, consts.HandleUtilTable)
				group.GET(consts.PathUtilForm, consts.HandleUtilForm)
				group.GET(consts.PathSituationMonitor, consts.HandleSituationMonitor)
				group.GET(consts.PathSituationOperation, consts.HandleSituationOperation)

			})
			s.Run()
			return nil
		},
	}
)
