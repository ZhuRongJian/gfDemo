// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfigData is the golang structure for table sys_config_data.
type SysConfigData struct {
	Id        uint64      `json:"id"         ` // 主键ID
	ConfigId  uint64      `json:"config_id"  ` // 配置ID
	Title     string      `json:"title"      ` // 配置标题
	Code      string      `json:"code"       ` // 配置编码
	Value     string      `json:"value"      ` // 配置值
	Status    uint        `json:"status"     ` // 状态：1正常 2停用
	Sort      uint        `json:"sort"       ` // 排序
	Note      string      `json:"note"       ` // 配置说明
	CreatedBy uint64      `json:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	UpdatedBy uint64      `json:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updated_at" ` // 更新时间
}
