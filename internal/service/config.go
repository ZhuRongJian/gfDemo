package service

import (
	"context"
	"errors"
	"gfDemo/internal/model"
	"gfDemo/internal/model/entity"
	"gfDemo/internal/service/internal/dao"
	"gfDemo/utility/query"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Config = new(serviceConfig)

type serviceConfig struct {
}

// ConfigQuery 配置值列表
func (s *serviceConfig) ConfigQuery(ctx context.Context, q *model.ConfigQuery) (*query.Result, error) {
	configQuery := make([]*entity.SysConfig, 0)
	return query.Page(dao.SysConfig.Ctx(ctx), q, &configQuery)
}


// ConfigChange 修改/新增配置值
func (s *serviceConfig) ConfigChange(ctx context.Context, vo *model.ConfigVo) error {
	cur := model.Context.Get(ctx)
	if vo.Id > 0 {
		temp := make([]model.ConfigVo, 0)
		err := dao.SysConfig.Ctx(ctx).Where("code=?", vo.Code).Scan(&temp)
		if err != nil {
			return err
		}
		for _, v := range temp {
			if v.Code == vo.Code {
				if v.Id != vo.Id {
					return errors.New("配置编码已存在")
				}
			}
		}
		vo.UpdatedBy = cur.Id
		_, err = dao.SysConfig.Ctx(ctx).WherePri(vo.Id).Data(vo).Update()
		if err != nil {
			return err
		}
	} else {
		if one, _ := dao.SysConfig.Ctx(ctx).One("code=?", vo.Code); !one.IsEmpty() {
			return errors.New("配置编码已存在")
		}
		if vo.CreatedBy == 0 {
			vo.CreatedBy = cur.Id
		}
		_, err := dao.SysConfig.Ctx(ctx).Data(vo).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

// ConfigDelete 删除配置值
func (s *serviceConfig) ConfigDelete(ctx context.Context, id uint) error {
	if one, _ := dao.SysConfigData.Ctx(ctx).One("config_id=?", id); !one.IsEmpty() {
		return errors.New("请先删除下级配置")
	}

	_, err := dao.SysConfig.Ctx(ctx).Where("id=?", id).Delete()
	return err
}

// ConfigDataQuery 获取配置类型列表
func (s *serviceConfig) ConfigDataQuery(ctx context.Context, q *model.ConfigDataQuery) (*query.Result, error) {
	var configDataQuery = make([]*entity.SysConfigData, 0)
	return query.Page(dao.SysConfigData.Ctx(ctx), q, &configDataQuery)
}

// ConfigDataChange 修改/新增配置类型
func (s *serviceConfig) ConfigDataChange(ctx context.Context, vo *model.ConfigDataVo) error {
	cur := model.Context.Get(ctx)
	if vo.Id > 0 {
		temp := make([]model.ConfigDataVo, 0)
		err := dao.SysConfigData.Ctx(ctx).Where("config_id=? and code=?", vo.ConfigId, vo.Code).Scan(&temp)
		if err != nil {
			return err
		}
		for _, v := range temp {
			if v.Code == vo.Code {
				if v.Id != vo.Id {
					return errors.New("配置编码已存在")
				}
			}
		}
		vo.UpdatedBy = cur.Id
		_, err = dao.SysConfigData.Ctx(ctx).WherePri(vo.Id).Data(vo).Update()
		if err != nil {
			return err
		}
	} else {
		if one, _ := dao.SysConfigData.Ctx(ctx).One("config_id=? and code=?", vo.ConfigId, vo.Code); !one.IsEmpty() {
			return errors.New("配置编码已存在")
		}
		if vo.CreatedBy == 0 {
			vo.CreatedBy = cur.Id
		}
		_, err := dao.SysConfigData.Ctx(ctx).Data(vo).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

// ConfigDataDelete 删除配置类型
func (s *serviceConfig) ConfigDataDelete(ctx context.Context, id uint) error {
	_, err := dao.SysConfigData.Ctx(ctx).Where("id=?", id).Delete()
	if err != nil {
		return gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	return err
}
