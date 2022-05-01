package v1

import (
	"gfDemo/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminQueryReq struct {
	g.Meta `path:"/admin/query" tags:"Admin" method:"get" summary:"获取用户列表"`
	*model.AdminQuery
}

type AdminQueryRes []*model.AdminListVO
