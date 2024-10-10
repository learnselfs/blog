// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tag

import (
	"context"

	"blog/api/tag/v1"
)

type ITagV1 interface {
	Tag(ctx context.Context, req *v1.TagReq) (res *v1.TagRes, err error)
}
