package v1

import (
	"gfDemo/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMenuListReq struct {
	g.Meta `path:"/menus" tags:"Index" method:"get" summary:"获取菜单列表"`
}

type GetMenuListRes []*model.MenuDTO


type UserInfoReq struct {
	g.Meta `path:"/index/getUserInfo" tags:"Index" method:"get" summary:"获取用户信息"`
}

type UserInfoRes struct {
	*model.AdminVO
}
