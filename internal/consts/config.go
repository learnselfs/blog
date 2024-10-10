// Package consts @Author Bing
// @Date 2024/9/6 10:05:00
// @Desc
package consts

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

var (
	PearConfig g.Map
)

func init() {
	PearConfig = GetPearConfig()
}

func GetPearConfig() g.Map {
	ctx := context.TODO()
	g.Cfg("pear").GetAdapter().(*gcfg.AdapterFile).SetFileName("pear.config.yml")
	config, err := g.Cfg("pear").Data(ctx)
	if err != nil {
		g.Log("config").Panic(ctx, "pear config err")
	}
	return config

}
