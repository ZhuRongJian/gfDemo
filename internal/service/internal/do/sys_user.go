// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta    `orm:"table:sys_user, do:true"`
	Id        interface{} // 主键
	OpenId    interface{} //
	UnionId   interface{} //
	Nickname  interface{} // 昵称
	Phone     interface{} // 手机
	Avatar    interface{} // 头像
	RealName  interface{} //
	CreatedAt *gtime.Time // 创建时间
	CreatedBy interface{} // 创建人
	UpdatedAt *gtime.Time // 更新时间
	UpdatedBy interface{} // 更新人
}
