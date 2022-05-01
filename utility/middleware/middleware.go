package middleware

import "github.com/gogf/gf/v2/net/ghttp"

var Middleware = new(middleware)

type middleware struct {
}

func (m *middleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
