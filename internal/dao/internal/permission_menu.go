// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionMenuDao is the data access object for table permission_menu.
type PermissionMenuDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns PermissionMenuColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionMenuColumns defines and stores column names for table permission_menu.
type PermissionMenuColumns struct {
	Id           string // 标识
	PermissionId string // 权限标识
	MenuId       string // 目录标识
	Type         string // 类型: 可选 read/write/null: 读/写/读写)
}

// permissionMenuColumns holds the columns for table permission_menu.
var permissionMenuColumns = PermissionMenuColumns{
	Id:           "id",
	PermissionId: "permission_id",
	MenuId:       "menu_id",
	Type:         "type",
}

// NewPermissionMenuDao creates and returns a new DAO object for table data access.
func NewPermissionMenuDao() *PermissionMenuDao {
	return &PermissionMenuDao{
		group:   "default",
		table:   "permission_menu",
		columns: permissionMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PermissionMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PermissionMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PermissionMenuDao) Columns() PermissionMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PermissionMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PermissionMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PermissionMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
