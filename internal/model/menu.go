package model

import (
	"gfDemo/internal/model/entity"
)

type MenuDTO struct {
	entity.SysMenu
	Checked  bool       `json:"checked" orm:"-"`
	Open     bool       `json:"open" orm:"-"`
	Children []*MenuDTO `json:"children" orm:"-"`
}

type MenuVO struct {
	Id         uint64 `json:"id"         ` // 主键ID
	Pid        uint64 `json:"pid"        ` // 父级ID
	Title      string `json:"title"      ` // 菜单标题
	Icon       string `json:"icon"       ` // 图标
	Path       string `json:"path"       ` // 菜单路径
	Component  string `json:"component"  ` // 菜单组件
	Target     string `json:"target"     ` // 目标
	Permission string `json:"permission" ` // 权限标识
	ApiMethod  string `json:"apiMethod"  ` // GET/POST
	ApiPath    string `json:"apiPath"    ` // Api 请求地址
	Type       uint   `json:"type"       ` // 类型：0菜单/1按钮
	Status     uint   `json:"status"     ` // 是否显示：0禁用/1正常
	Note       string `json:"note"       ` // 备注
	Sort       uint   `json:"sort"       ` // 显示顺序
	CreatedBy  uint64 `json:"createdBy"  ` // 添加人
	UpdatedBy  uint64 `json:"updatedBy"  ` // 更新人
}
