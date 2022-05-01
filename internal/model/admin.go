package model

import (
	"gfDemo/utility/query"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

type AdminLoginVO struct {
	Account  string
	Password string
}

type AdminVO struct {
	Id          uint     `json:"id"        ` // 主键ID
	Nickname    string   `json:"nickname"`   // 昵称
	Mobile      string   `json:"mobile"`     // 手机号码
	Account     string   `json:"account"`    // 登录用户名
	Permissions []string `json:"permissionList"`
}

type AdminQuery struct {
	query.PageParam
}

func (q *AdminQuery) Cond(m *gdb.Model) *gdb.Model {
	return m
}

type  AdminListVO struct {
	Id        uint64      `json:"id"         ` // 主键ID
	RoleId    uint64      `json:"role_id"    ` // 角色ID
	Partners  string      `json:"partners"   ` //
	Nickname  string      `json:"nickname"   ` // 昵称
	Avatar    string      `json:"avatar"     ` // 头像
	Mobile    string      `json:"mobile"     ` // 手机号码
	Account   string      `json:"account"    ` //
	Password  string      `json:"password"   ` // 登录密码
	Salt      string      `json:"salt"       ` // 盐加密
	Status    uint        `json:"status"     ` //
	Note      string      `json:"note"       ` // 备注
	LoginNum  uint        `json:"login_num"  ` // 登录次数
	LoginIp   string      `json:"login_ip"   ` // 最近登录IP
	LoginTime *gtime.Time `json:"login_time" ` // 最近登录时间
	CreatedBy uint64      `json:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	UpdatedBy uint64      `json:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updated_at" ` // 更新时间
	RoleName  string      `json:"role_name"`
}