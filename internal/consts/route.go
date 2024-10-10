package consts

import "github.com/gogf/gf/v2/net/ghttp"

const (
	// admin prefix
	AdminPrefix = "/view"

	PathLogin                 = "GET:/login"
	PathIndex                 = "/index"
	PathAnalysisIndex         = "/analysis/index"
	PathComponentAdmin        = "/component/admin"
	PathComponentGrid         = "/component/grid"
	PathComponentNprogress    = "/component/nprogress"
	PathComponentToast        = "/component/toast"
	PathConsoleIndex          = "/console/index"
	PathException403          = "/exception/403"
	PathException404          = "/exception/404"
	PathException500          = "/exception/500"
	PathListingTable          = "/listing/table"
	PathProfileIndex          = "/profile/index"
	PathResultError           = "/result/error"
	PathResultSuccess         = "/result/success"
	PathRoleBaseAccessControl = "/system/manager"
	PathSystemMenu            = "/system/menu"
	PathSystemUser            = "/system/user"
	PathSystemRole            = "/system/role"
	PathSystemPermission      = "/system/permission"
	PathUtilTable             = "/util/table"
	PathUtilForm              = "/util/form"
	PathSituationMonitor      = "situation/monitor"
	PathSituationOperation    = "situation/operation"

	tplLogin                 = "./login.html"
	tplIndex                 = "./index.html"
	tplAnalysisIndex         = "./view/analysis/index.html"
	tplComponentAdmin        = "./view/component/admin.html"
	tplComponentGrid         = "./view/component/grid.html"
	tplComponentNprogress    = "./view/component/nprogress.html"
	tplComponentToast        = "./view/component/toast.html"
	tplConsoleIndex          = "./view/console/index.html"
	tplException403          = "./view/exception/403.html"
	tplException404          = "./view/exception/404.html"
	tplException500          = "./view/exception/500.html"
	tplListingTable          = "./view/listing/table.html"
	tplProfileIndex          = "./view/profile/index.html"
	tplResultError           = "./view/result/error.html"
	tplResultSuccess         = "./view/result/success.html"
	tplRoleBaseAccessControl = "./view/system/manager.html"
	tplSystemMenu            = "./view/system/menu.html"
	tplSystemUser            = "./view/system/user.html"
	tplSystemRole            = "./view/system/role.html"
	tplSystemPermission      = "./view/system/permission.html"
	tplUtilTable             = "./view/util/table.html"
	tplUtilForm              = "./view/util/Form.html"
	tplSituationMonitor      = "./view/situation/monitor.html"
	tplSituationOperation    = "./view/situation/operation.html"
)

var (
	HandleLogin = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplLogin); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleIndex = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplIndex); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleAnalysisIndex = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplAnalysisIndex); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleComponentAdmin = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplComponentAdmin); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleComponentGrid = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplComponentGrid); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleComponentNprogress = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplComponentNprogress); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleComponentToast = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplComponentToast); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleConsoleIndex = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplConsoleIndex); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleException403 = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplException403); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleException404 = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplException404); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleException500 = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplException500); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleListingTable = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplListingTable); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleProfileIndex = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplProfileIndex); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleResultError = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplResultError); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleResultSuccess = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplResultSuccess); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleRoleBaseAccessControl = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplRoleBaseAccessControl); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleSystemMenu = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplSystemMenu); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleSystemUser = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplSystemUser); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleSystemRole = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplSystemRole); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleSystemPermission = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplSystemPermission); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleUtilTable = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplUtilTable); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleUtilForm = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplUtilForm); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleSituationMonitor = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplSituationMonitor); err != nil {
			g.Response.Writeln(err)
		}
	}
	HandleSituationOperation = func(g *ghttp.Request) {
		if err := g.Response.WriteTpl(tplSituationOperation); err != nil {
			g.Response.Writeln(err)
		}
	}
)
