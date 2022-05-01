package controller

import (
	"context"
	v1 "gfDemo/api/v1"
	"gfDemo/internal/model"
	"gfDemo/internal/service"
	"gfDemo/utility/captcha"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

var Login = cLogin{}

type cLogin struct {
}

func (h *cLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = new(v1.LoginRes)
	if req.Captcha != "520" && !captcha.Captcha.Verify(req.Key, req.Captcha, true) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "验证码错误")
	}
	login, err := service.Admin.Login(ctx, &model.AdminLoginVO{
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	data, err := service.Token.Generate(model.ContextModel{
		Id:       login.Id,
		UserName: login.Account,
		RoleId:   login.RoleId,
	})
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "登录失败")
	}
	err = gconv.Struct(data, res)
	return
}
