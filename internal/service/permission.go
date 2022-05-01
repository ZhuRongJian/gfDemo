package service

import (
	"context"
	"fmt"
	"gfDemo/internal/model"
	"gfDemo/internal/model/entity"
	"gfDemo/internal/service/internal/dao"
	"gfDemo/utility/permission"
	"gfDemo/utility/response"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var Permission = new(servicePermission)

type servicePermission struct {
}

func (s *servicePermission) Middleware(r *ghttp.Request) {
	path := r.URL.Path
	method := r.Method
	cur := model.Context.Get(r.GetCtx())
	if s.CheckPerm(r.GetCtx(), cur.RoleId, path, method) {
		r.Middleware.Next()
	} else {
		_ = r.Response.WriteJsonExit(response.Error(gerror.NewCode(gcode.CodeBusinessValidationFailed, "暂无权限操作此功能")))
		return
	}
}
func (s *servicePermission) CheckPerm(ctx context.Context, roleId uint64, path, method string) bool {
Loop:
	key := fmt.Sprintf("role:%d", roleId)
	v, err := g.DB().GetCache().Get(ctx, key)
	if err == nil && !v.IsNil() {
		m := v.Array()
		for _, v := range m {
			var menu entity.SysMenu
			if err := gconv.Struct(v, &menu); err == nil {
				if permission.KeyMatch(path, menu.ApiPath) && menu.ApiMethod == method {
					return true
				}
			}
		}
		return false
	}
	if err := s.LoadPolicy(ctx, roleId); err != nil {
		return false
	}
	goto Loop
}

func (s *servicePermission) LoadPolicy(ctx context.Context, roleId uint64) error {
	var menus []entity.SysMenu
	err := dao.SysRoleMenu.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     fmt.Sprintf("role:%d", roleId),
	}).
		InnerJoin(dao.SysMenu.Table(), "sys_role_menu.menu_id=sys_menu.id").
		Where("sys_role_menu.role_id=? and sys_menu.type=1", roleId).
		Scan(&menus)
	return err
}

func (s *servicePermission) FlushCache(ctx context.Context) error {
	var roles []*entity.SysRole
	if err := dao.SysRole.Ctx(ctx).Scan(&roles); err != nil {
		return err
	}

	var keys []interface{}
	for _, v := range roles {
		keys = append(keys, fmt.Sprintf("role:%d", v.Id))
	}
	_, err := g.DB().GetCache().Remove(ctx, keys...)
	return err
}
