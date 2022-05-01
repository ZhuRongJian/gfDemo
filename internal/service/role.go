package service

import (
	"context"
	"gfDemo/internal/model"
	"gfDemo/internal/model/entity"
	"gfDemo/internal/service/internal/dao"
	"gfDemo/utility/query"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Role = new(serviceRole)

type serviceRole struct {
}

func (s *serviceRole) RoleQuery(ctx context.Context, q *model.RoleQuery) (res []*entity.SysRole, err error) {
	err = query.Exec(dao.SysRole.Ctx(ctx), q, &res)
	return
}

func (s *serviceRole) RoleChange(ctx context.Context, vo *model.RoleVo) (err error) {
	cur := model.Context.Get(ctx)
	if vo.Id > 0 {
		vo.UpdatedBy = cur.Id
		//修改逻辑
		_, err = dao.SysRole.Ctx(ctx).WherePri(vo.Id).Data(vo).Update()
	} else {
		vo.CreatedBy = cur.Id
		_, err = dao.SysRole.Ctx(ctx).Insert(vo)
	}
	return nil
}

func (s *serviceRole) RoleDelete(ctx context.Context, id uint) (err error) {
	cnt, err := dao.SysAdmin.Ctx(ctx).Where("role_id=?", id).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.NewCode(gcode.CodeValidationFailed, "此角色正在被使用")
	}
	_, err = dao.SysRole.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	return nil
}

func (s *serviceRole) SavePermission(ctx context.Context, id uint64, menuIds []uint64) (err error) {
	var rms []*entity.SysRoleMenu
	for _, v := range menuIds {
		rms = append(rms, &entity.SysRoleMenu{
			RoleId: id,
			MenuId: v,
		})
	}
	//defer func() {
	//	_ = Permission.FlushCache(ctx)
	//}()
	return dao.SysRoleMenu.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if _, err := dao.SysRoleMenu.Ctx(ctx).Delete("role_id=?", id); err != nil {
			return gerror.NewCode(gcode.CodeDbOperationError, err.Error())
		}
		if len(rms) > 0 {
			if _, err := dao.SysRoleMenu.Ctx(ctx).Insert(&rms); err != nil {
				return gerror.NewCode(gcode.CodeDbOperationError, err.Error())
			}
		}
		return nil
	})
}
