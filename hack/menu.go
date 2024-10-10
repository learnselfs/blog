// Package hack @Author Bing
// @Date 2024/9/6 12:53:00
// @Desc
package hack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

var res = bytes.NewBufferString("")
var ctx = context.TODO()
var in = make([]any, 0)

func MenuLoad() {
	s, _ := load()

	for i, v := range s {
		v1, ok := v.(g.Map)
		if ok {
			in = append(in, v1)
		}
		var v2 g.Map
		json.Unmarshal([]byte(i), &v2)
		in = append(in, v2)
	}
	menuCreate(in)
}

func menuCreate(s []any, id ...any) {
	for _, v1 := range s {
		v, ok := v1.(g.Map)
		if !ok {
			continue
		}
		if len(id) > 0 {
			v["parentMid"] = id[0]
		}

		res.Reset()
		form, _ := g.Client().Post(ctx, "localhost:8000/menu", v)
		res.ReadFrom(form.Response.Body)
		var r g.Map
		var mid any
		json.Unmarshal(res.Bytes(), &r)
		data, ok := r["data"].(g.Map)
		if ok {
			menu, ok := data["menu"].(g.Map)
			if ok {
				mid = menu["mid"]
			}
		}
		if v["children"] != nil {
			value, ok := v["children"].([]any)
			if ok {
				menuCreate(value, mid)
			}
		}
		fmt.Println(res, mid)
	}
}

func load() (g.Map, error) {
	menu := "menu"
	g.Cfg(menu).GetAdapter().(*gcfg.AdapterFile).SetFileName("menu.json")
	return g.Cfg(menu).Data(ctx)

}
