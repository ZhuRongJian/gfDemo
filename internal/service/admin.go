package service

import (
	"context"
	"gfDemo/internal/model"
	"gfDemo/internal/model/entity"
	"gfDemo/internal/service/internal/dao"
	"gfDemo/utility/password"
	"gfDemo/utility/query"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Admin = new(serviceAdmin)
)

type serviceAdmin struct {
}

func (s *serviceAdmin) Login(ctx context.Context, req *model.AdminLoginVO) (*entity.SysAdmin, error) {
	r := g.RequestFromCtx(ctx)
	var admin entity.SysAdmin
	if err := dao.SysAdmin.Ctx(ctx).Where("account=?", req.Account).Scan(&admin); err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	if admin.Status == 2 {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "该账号已经被禁止登录")
	}
	if password.Verify(admin.Password, req.Password, admin.Salt) {
		//更新登录信息
		_, _ = dao.SysAdmin.Ctx(ctx).WherePri(admin.Id).Data(g.Map{
			dao.SysAdmin.Columns().LoginIp:   r.GetClientIp(),
			dao.SysAdmin.Columns().LoginNum:  gdb.Raw(dao.SysAdmin.Columns().LoginNum + "+1"),
			dao.SysAdmin.Columns().LoginTime: gtime.Now(),
		}).Unscoped().Update()
		return &admin, nil
	}

	return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "账号或密码错误")
}

func (s *serviceAdmin) UserInfo(ctx context.Context) (*model.AdminVO, error) {
	cur := model.Context.Get(ctx)
	var admin entity.SysAdmin
	if err := dao.SysAdmin.Ctx(ctx).Where("id=?", cur.Id).Scan(&admin); err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	var v model.AdminVO
	_ = gconv.Struct(&admin, &v)
	ps, err := Menu.GetPermissions(ctx, cur.Id)
	if err != nil {
		return nil, err
	}
	v.Permissions = ps
	return &v, nil
}

func (s *serviceAdmin) AdminQuery(ctx context.Context, q *model.AdminQuery) (res []*model.AdminListVO, err error) {

	m := dao.SysAdmin.Ctx(ctx).
		InnerJoin("sys_role", "sys_role.id=sys_admin.role_id")

	err = query.Exec(m, q, &res,"sys_admin.*,sys_role.name role_name")
	return
}
