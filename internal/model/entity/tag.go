// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure for table tag.
type Tag struct {
	Id         int         `json:"id"         orm:"id"          ` // Primary Key
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" ` // Create Time
	Name       string      `json:"name"       orm:"name"        ` // Tag Name
	Department string      `json:"department" orm:"department"  ` // tag department
}
