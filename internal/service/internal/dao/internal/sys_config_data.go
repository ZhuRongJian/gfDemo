// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigDataDao is the data access object for table sys_config_data.
type SysConfigDataDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysConfigDataColumns // columns contains all the column names of Table for convenient usage.
}

// SysConfigDataColumns defines and stores column names for table sys_config_data.
type SysConfigDataColumns struct {
	Id        string // 主键ID
	ConfigId  string // 配置ID
	Title     string // 配置标题
	Code      string // 配置编码
	Value     string // 配置值
	Status    string // 状态：1正常 2停用
	Sort      string // 排序
	Note      string // 配置说明
	CreatedBy string // 添加人
	CreatedAt string // 创建时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
}

//  sysConfigDataColumns holds the columns for table sys_config_data.
var sysConfigDataColumns = SysConfigDataColumns{
	Id:        "id",
	ConfigId:  "config_id",
	Title:     "title",
	Code:      "code",
	Value:     "value",
	Status:    "status",
	Sort:      "sort",
	Note:      "note",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewSysConfigDataDao creates and returns a new DAO object for table data access.
func NewSysConfigDataDao() *SysConfigDataDao {
	return &SysConfigDataDao{
		group:   "default",
		table:   "sys_config_data",
		columns: sysConfigDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysConfigDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysConfigDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysConfigDataDao) Columns() SysConfigDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysConfigDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysConfigDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysConfigDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}