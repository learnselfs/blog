package hello

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Hello(res *ghttp.Request) {
	res.Response.Writeln("hello !!!!!")
	return
}
