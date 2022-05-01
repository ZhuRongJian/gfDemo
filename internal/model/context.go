package model

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	//ContextKey 上下文变量存储键名，前后端系统共享
	ContextKey = "ContextKey"
)

//ContextModel 请求上下文结构
type ContextModel struct {
	Id       uint64 // 用户ID
	RoleId   uint64
	UserName string // 用户名称
	Platform string
}

// 上下文管理服务
var Context = contextShared{}

type contextShared struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextShared) Init(r *ghttp.Request) *ContextModel {
	value := r.GetCtxVar(ContextKey)
	model := new(ContextModel)
	if value.IsEmpty() {
		r.SetCtxVar(ContextKey, model)
		return model
	}
	_ = value.Scan(model)
	return model
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextShared) Get(ctx context.Context) *ContextModel {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*ContextModel); ok {
		return localCtx
	}
	return nil
}
