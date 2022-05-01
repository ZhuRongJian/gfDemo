package controller

import (
	"context"
	v1 "gfDemo/api/v1"
	"gfDemo/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Menu = cMenu{}

type cMenu struct {
}

func (h *cMenu) MenuList(ctx context.Context, req *v1.GetMenuIndexReq) (res v1.GetMenuListRes, err error) {
	if data, err := service.Menu.GetMenuList(ctx); err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	} else {
		return data, nil
	}
}

func (h *cMenu) MenuChange(ctx context.Context, req *v1.MenuChangeReq) (*v1.BaseEmptyRes, error) {
	return nil, service.Menu.MenuChange(ctx, req.MenuVO)
}

func (h *cMenu) MenuDelete(ctx context.Context, req *v1.MenuDeleteReq) (*v1.BaseEmptyRes, error) {
	return nil, service.Menu.MenuDelete(ctx, req.Id)
}
