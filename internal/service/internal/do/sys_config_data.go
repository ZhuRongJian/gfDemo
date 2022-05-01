// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfigData is the golang structure of table sys_config_data for DAO operations like Where/Data.
type SysConfigData struct {
	g.Meta    `orm:"table:sys_config_data, do:true"`
	Id        interface{} // 主键ID
	ConfigId  interface{} // 配置ID
	Title     interface{} // 配置标题
	Code      interface{} // 配置编码
	Value     interface{} // 配置值
	Status    interface{} // 状态：1正常 2停用
	Sort      interface{} // 排序
	Note      interface{} // 配置说明
	CreatedBy interface{} // 添加人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy interface{} // 更新人
	UpdatedAt *gtime.Time // 更新时间
}
