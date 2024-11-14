// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChaoxingUserDao is the data access object for table chaoxing_user.
type ChaoxingUserDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ChaoxingUserColumns // columns contains all the column names of Table for convenient usage.
}

// ChaoxingUserColumns defines and stores column names for table chaoxing_user.
type ChaoxingUserColumns struct {
	Id         string //
	Phone      string // 超星登录手机号
	Password   string // 登录密码
	Cookies    string // 登录cookies
	UpdateTime string //
	CreateTime string //
}

// chaoxingUserColumns holds the columns for table chaoxing_user.
var chaoxingUserColumns = ChaoxingUserColumns{
	Id:         "id",
	Phone:      "phone",
	Password:   "password",
	Cookies:    "cookies",
	UpdateTime: "update_time",
	CreateTime: "create_time",
}

// NewChaoxingUserDao creates and returns a new DAO object for table data access.
func NewChaoxingUserDao() *ChaoxingUserDao {
	return &ChaoxingUserDao{
		group:   "default",
		table:   "chaoxing_user",
		columns: chaoxingUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChaoxingUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChaoxingUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChaoxingUserDao) Columns() ChaoxingUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChaoxingUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChaoxingUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChaoxingUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
