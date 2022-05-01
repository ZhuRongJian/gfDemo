package v1

import (
	"gfDemo/internal/model"
	"gfDemo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleQueryReq struct {
	g.Meta `path:"/role/query" tags:"Role" method:"get" summary:"获取角色列表"`
	*model.RoleQuery
}

type RoleQueryRes []*entity.SysRole

type RoleChangeReq struct {
	g.Meta `path:"/role/change" tags:"Role" method:"post" summary:"添加角色"`
	*model.RoleVo
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete/:id" tags:"Role" method:"post" summary:"删除角色"`
	Id     uint
}
type RolePermissionReq struct {
	g.Meta `path:"/role/premission/:id" tags:"Role" method:"get" summary:"角色权限"`
	Id     uint
}

type RolePermissionRes []*model.MenuDTO

type SavePermissionReq struct {
	g.Meta  `path:"/role/save_permission" tags:"Role" method:"post" summary:"设置角色权限"`
	RoleId  uint64 `v:"required#请选择角色"`
	MenuIds []uint64
}
