package user

import (
	v1 "blog/api/user/v1"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) AddUser(ctx context.Context, req *v1.PostReq) (res *v1.PostRes, err error) {

	return
}

func (c *Controller) UpdateUser(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return
}

func (c *Controller) DeleteUser(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return
}

func (c *Controller) ListUser(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	return
}

func (c *Controller) GetUser(ctx context.Context, req *v1.OneReq) (res *v1.OneRes, err error) {

	return
}

func (c *Controller) Post(r *ghttp.Request) {
	r.Response.Writeln("添加用户")
}

func (c *Controller) Put(r *ghttp.Request) {
	r.Response.Writeln("更新用户")
}

func (c *Controller) Delete(r *ghttp.Request) {
	r.Response.Writeln("删除用户")
}

func (c *Controller) Get(r *ghttp.Request) {
	r.Response.Writeln("查询一个用户")
}

// 参数
func (c *Controller) Params(request *ghttp.Request) {
	m := request.GetQueryMap()
	request.Response.WriteJson(m)
}

func (c *Controller) Route(ctx context.Context, req *v1.OneReq) (res *v1.OneRes, err error) {
	// request := g.RequestFromCtx(ctx)
	// // u := request.GetRouter("name")
	// queryMap := request.GetQueryMap(g.Map{"name":"guest", "age": 18})
	// var p v1.

	// p.Age, err = strconv.Atoi(queryMap["age"].(string))
	// p.Name = queryMap["name"].(string)
	// res = &p
	return
}
