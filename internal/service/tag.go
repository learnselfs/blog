// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ITag interface {
		GetList(ctx context.Context)
	}
)

var (
	localTag ITag
)

func Tag() ITag {
	if localTag == nil {
		panic("implement not found for interface ITag, forgot register?")
	}
	return localTag
}

func RegisterTag(i ITag) {
	localTag = i
}
