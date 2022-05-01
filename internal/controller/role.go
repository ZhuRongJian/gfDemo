package controller

import (
	"context"
	v1 "gfDemo/api/v1"
	"gfDemo/internal/service"
)

var Role = cRole{}

type cRole struct {
}

func (h *cRole) RoleQuery(ctx context.Context, req *v1.RoleQueryReq) (v1.RoleQueryRes, error) {
	return service.Role.RoleQuery(ctx, req.RoleQuery)
}
func (h *cRole) RoleChange(ctx context.Context, req *v1.RoleChangeReq) (*v1.BaseEmptyRes, error) {
	return nil, service.Role.RoleChange(ctx, req.RoleVo)
}
func (h *cRole) RoleDelete(ctx context.Context, req *v1.RoleDeleteReq) (*v1.BaseEmptyRes, error) {
	return nil, service.Role.RoleDelete(ctx, req.Id)
}
func (h *cRole) Permission(ctx context.Context, req *v1.RolePermissionReq) (v1.RolePermissionRes, error) {
	return service.Menu.GetMenuByRole(ctx, req.Id)
}
func (h *cRole) SavePermission(ctx context.Context, req *v1.SavePermissionReq) (*v1.BaseEmptyRes, error) {
	return nil, service.Role.SavePermission(ctx, req.RoleId, req.MenuIds)
}
