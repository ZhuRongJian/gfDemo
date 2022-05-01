package service

import (
	"context"
	"gfDemo/internal/model"
	"gfDemo/utility/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
	"time"
)

// CustomClaims 自定义的 metadata在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	Data model.ContextModel
	jwt.StandardClaims
}

var Token = newToken()

// token token
type token struct {
	Secret string
	Expire int
}

//Setup 创建中间件
func newToken() *token {
	var t token
	v, err := g.Cfg().Get(context.Background(), "jwt")
	if err != nil {
		log.Fatalln(err)
	}
	err = v.Scan(&t)
	if err != nil {
		log.Fatalln(err)
	}
	return &t
}
func (t *token) Middleware(r *ghttp.Request) {
	token := r.Request.Header.Get("Authorization")
	if token == "" {
		r.Response.WriteJsonExit(response.Error(gerror.NewCode(gcode.CodeNotAuthorized, "登录状态过期,请重新登录!")))
		return
	}
	data, err := t.Valid(token)
	if err != nil {
		r.Response.WriteJsonExit(response.Error(gerror.NewCode(gcode.CodeNotAuthorized, "登录状态过期,请重新登录!")))
		return
	}
	res := model.Context.Init(r)
	res.Id = data.Data.Id
	res.UserName = data.Data.UserName
	res.RoleId = data.Data.RoleId
	res.Platform = r.Request.Header.Get("X-Access-Platform")
	r.Middleware.Next()
}

//Valid 解码
func (t *token) Valid(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	// 解密转换类型并返回
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Generate 将 User 用户信息加密为 JWT 字符串
// expireTime := time.Now().Add(time.Hour * 24 * 3).Unix() 三天后过期
func (t *token) Generate(data model.ContextModel) (g.Map, error) {
	claims := CustomClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			Issuer:   "clean-sass",
			IssuedAt: time.Now().Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := jwtToken.SignedString([]byte(t.Secret))
	if err != nil {
		return nil, err
	}
	return g.Map{
		"access_token": s,
		"token_type":   "Bearer",
	}, nil
}
