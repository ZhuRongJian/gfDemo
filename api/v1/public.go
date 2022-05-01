package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaReq struct {
	g.Meta `path:"/captcha" tags:"Base" method:"get" summary:"验证码"`
}

type CaptchaRes struct {
	Key     string `json:"key"`
	Captcha string `json:"captcha"`
}

