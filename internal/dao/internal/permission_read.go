// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionReadDao is the data access object for table permission_read.
type PermissionReadDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns PermissionReadColumns // columns contains all the column names of Table for convenient usage.
}

// PermissionReadColumns defines and stores column names for table permission_read.
type PermissionReadColumns struct {
	Id           string // 标识
	PermissionId string // 权限标识
	MenuId       string // 目录标识
	Type         string // 类型: 可选 read/write/null: 读/写/读写)
}

// permissionReadColumns holds the columns for table permission_read.
var permissionReadColumns = PermissionReadColumns{
	Id:           "id",
	PermissionId: "permission_id",
	MenuId:       "menu_id",
	Type:         "type",
}

// NewPermissionReadDao creates and returns a new DAO object for table data access.
func NewPermissionReadDao() *PermissionReadDao {
	return &PermissionReadDao{
		group:   "default",
		table:   "permission_read",
		columns: permissionReadColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PermissionReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PermissionReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PermissionReadDao) Columns() PermissionReadColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PermissionReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PermissionReadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PermissionReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
