package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Index" method:"post" summary:"登录接口"`
	Account string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
	Key      string `json:"key"`
	Remember bool   `json:"remember"`
}
type LoginRes struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

