// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManyRolePermissionDao is the data access object for table many_role_permission.
type ManyRolePermissionDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ManyRolePermissionColumns // columns contains all the column names of Table for convenient usage.
}

// ManyRolePermissionColumns defines and stores column names for table many_role_permission.
type ManyRolePermissionColumns struct {
	Id  string // id
	Rid string // 角色 id
	Pid string // 权限 id
}

// manyRolePermissionColumns holds the columns for table many_role_permission.
var manyRolePermissionColumns = ManyRolePermissionColumns{
	Id:  "id",
	Rid: "rid",
	Pid: "pid",
}

// NewManyRolePermissionDao creates and returns a new DAO object for table data access.
func NewManyRolePermissionDao() *ManyRolePermissionDao {
	return &ManyRolePermissionDao{
		group:   "default",
		table:   "many_role_permission",
		columns: manyRolePermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ManyRolePermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ManyRolePermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ManyRolePermissionDao) Columns() ManyRolePermissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ManyRolePermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ManyRolePermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ManyRolePermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
