package controller

import (
	"context"
	v1 "gfDemo/api/v1"
	"gfDemo/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Config = new(cConfig)

type cConfig struct {
}

// ConfigQuery 查询配置值
func (h *cConfig) ConfigQuery(ctx context.Context, req *v1.ConfigQueryReq) (*v1.BasePageQueryRes, error) {
	if r, err := service.Config.ConfigQuery(ctx, req.ConfigQuery); err != nil {
		return nil, err
	} else {
		return &v1.BasePageQueryRes{Result: r}, nil
	}
}

// ConfigChange 修改/新增字典值
func (h *cConfig) ConfigChange(ctx context.Context, req *v1.ConfigChangeReq) (*v1.BaseEmptyRes, error) {
	if err := service.Config.ConfigChange(ctx, req.ConfigVo); err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	return nil, nil
}

// ConfigDelete 删除字典值
func (h *cConfig) ConfigDelete(ctx context.Context, req *v1.ConfigDeleteReq) (*v1.BaseEmptyRes, error) {
	return nil, service.Config.ConfigDelete(ctx, req.Id)
}


// ConfigDataQuery 查询配置类型
func (h *cConfig) ConfigDataQuery(ctx context.Context, req *v1.ConfigDataQueryReq) (*v1.BasePageQueryRes, error) {
	if r, err := service.Config.ConfigDataQuery(ctx, req.ConfigDataQuery); err != nil {
		return nil, err
	} else {
		return &v1.BasePageQueryRes{Result: r}, nil
	}
}

// ConfigDataChange 修改/新增配置类型
func (h *cConfig) ConfigDataChange(ctx context.Context, req *v1.ConfigDataChangeReq) (*v1.BaseEmptyRes, error) {
	if err := service.Config.ConfigDataChange(ctx, req.ConfigDataVo); err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	return nil, nil
}

// ConfigDataDelete 删除配置类型
func (h *cConfig) ConfigDataDelete(ctx context.Context, req *v1.ConfigDataDeleteReq) (*v1.BaseEmptyRes, error) {
	return nil, service.Config.ConfigDataDelete(ctx, req.Id)
}