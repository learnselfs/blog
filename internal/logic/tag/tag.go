// Package tag @Author Bing
// @Date 2024/8/27 15:40:00
// @Desc
package tag

import (
	"blog/internal/service"
	"context"
)

func init() {
	service.RegisterTag(New())
}

type sTag struct{}

func New() *sTag { return &sTag{} }

func (s *sTag) GetList(ctx context.Context) {

}
