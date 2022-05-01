package service

import (
	"context"
	"gfDemo/internal/model"
	"gfDemo/internal/service/internal/dao"
)

var Index = new(serviceIndex)

type serviceIndex struct {
}

func (s *serviceIndex) GetMenuList(ctx context.Context) ([]*model.MenuDTO, error) {
	m := model.Context.Get(ctx)
	if m.Id == 1 {
		return dao.GetMenuAllTree(ctx)
	} else {
		return dao.GetMenuAllTreeByRoleId(ctx, m.RoleId)
	}
}
