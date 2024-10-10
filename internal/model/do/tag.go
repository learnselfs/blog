// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure of table tag for DAO operations like Where/Data.
type Tag struct {
	g.Meta     `orm:"table:tag, do:true"`
	Id         interface{} // Primary Key
	CreateTime *gtime.Time // Create Time
	Name       interface{} // Tag Name
	Department interface{} // tag department
}
