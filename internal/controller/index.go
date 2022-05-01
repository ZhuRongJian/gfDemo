package controller

import (
	"context"
	v1 "gfDemo/api/v1"
	"gfDemo/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Index = cIndex{}

type cIndex struct {
}

func (h *cIndex) MenuList(ctx context.Context, req *v1.GetMenuListReq) (res v1.GetMenuListRes, err error) {
	if data, err := service.Index.GetMenuList(ctx); err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	} else {
		return data, nil
	}
}
func (h *cIndex) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	if data, err := service.Admin.UserInfo(ctx); err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	} else {
		return &v1.UserInfoRes{AdminVO: data}, nil
	}
}
