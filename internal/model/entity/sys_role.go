// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id        uint64      `json:"id"         ` //
	Name      string      `json:"name"       ` //
	Code      string      `json:"code"       ` // 角色编码
	Sort      int         `json:"sort"       ` //
	Note      string      `json:"note"       ` // 备注
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	CreatedBy uint64      `json:"created_by" ` // 创建人
	UpdatedAt *gtime.Time `json:"updated_at" ` // 更新时间
	UpdatedBy uint64      `json:"updated_by" ` // 更新人
}
