package controller

import (
	"context"
	v1 "gfDemo/api/v1"
	"gfDemo/utility/captcha"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Public = cPublic{}

type cPublic struct {
}

func (h *cPublic) GenCode(ctx context.Context, req *v1.CaptchaReq) (*v1.CaptchaRes, error) {
	key, base, err := captcha.Captcha.Generate()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	return &v1.CaptchaRes{
		Key:     key,
		Captcha: base,
	}, err
}
