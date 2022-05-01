package service

import (
	"context"
	"gfDemo/internal/model"
	"gfDemo/internal/model/entity"
	"gfDemo/internal/service/internal/dao"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Menu = new(serviceMenu)

type serviceMenu struct {
}

func (s *serviceMenu) GetMenuList(ctx context.Context) ([]*model.MenuDTO, error) {
	m := model.Context.Get(ctx)
	if m.Id == 1 {
		return dao.GetMenuAll(ctx)
	} else {
		return dao.GetMenuAllTreeByRoleId(ctx, m.RoleId)
	}
}

func (s *serviceMenu) GetPermissions(ctx context.Context, userId uint64) (permissions []string, err error) {
	var (
		menus []entity.SysMenu
	)
	if userId == 1 {
		// 超级管理员
		if err = dao.SysMenu.Ctx(ctx).Where("type=1 and status=1").Scan(&menus); err != nil {
			return
		}
	} else {
		// 其他
		menus, err = dao.GetPermissions(ctx, userId)
		if err != nil {
			return
		}
	}
	for _, v := range menus {
		permissions = append(permissions, v.Permission)
	}
	return permissions, nil
}

func (s *serviceMenu) MenuChange(ctx context.Context, vo *model.MenuVO) (err error) {
	cur := model.Context.Get(ctx)

	if vo.Id > 0 {
		//defer func() {
		//	_ = Permission.FlushCache(ctx)
		//}()

		vo.UpdatedBy = cur.Id
		//修改逻辑
		_, err = dao.SysMenu.Ctx(ctx).WherePri(vo.Id).Data(vo).Update()
	} else {
		vo.CreatedBy = cur.Id
		_, err = dao.SysMenu.Ctx(ctx).Insert(vo)
	}
	return
}

func (s *serviceMenu) MenuDelete(ctx context.Context, id uint) (err error) {
	cnt, err := dao.SysMenu.Ctx(ctx).Where("pid=?", id).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.NewCode(gcode.CodeValidationFailed, "请先删除所有子节点")
	}

	err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if _, err := dao.SysRoleMenu.Ctx(ctx).Delete("menu_id=?", id); err != nil {
			return err
		}
		_, err := dao.SysMenu.Ctx(ctx).Delete("id=?", id)
		return err
	})
	if err != nil {
		return gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	return nil
}

func (s *serviceMenu) GetMenuByRole(ctx context.Context, roleId uint) ([]*model.MenuDTO, error) {
	menus, _ := dao.GetMenuAll(ctx)
	if len(menus) > 0 {
		var temp []*entity.SysRoleMenu
		if err := dao.SysRoleMenu.Ctx(ctx).Where("role_id=?", roleId).Scan(&temp); err != nil {
			return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
		}
		for _, v := range menus {
			for _, v1 := range temp {
				if v.Id == v1.MenuId {
					v.Checked = true
					v.Open = true
					break
				}
			}
		}
	}
	return menus, nil
}
