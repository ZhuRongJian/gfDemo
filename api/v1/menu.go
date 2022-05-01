package v1

import (
	"gfDemo/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMenuIndexReq struct {
	g.Meta `path:"/menu/index" tags:"Index" method:"get" summary:"获取菜单列表"`
}

type MenuChangeReq struct {
	g.Meta `path:"/menu/change" tags:"Index" method:"post" summary:"菜单修改"`
	*model.MenuVO
}

type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete" tags:"Menu" method:"post" summary:"删除菜单"`
	Id     uint `v:"required#请选择要删除的"`
}

