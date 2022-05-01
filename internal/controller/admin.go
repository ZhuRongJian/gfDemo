package controller

import (
	"context"
	v1 "gfDemo/api/v1"
	"gfDemo/internal/service"
)

var Admin = new(cAdmin)

type cAdmin struct {
}

func (h *cAdmin) AdminQuery(ctx context.Context, req *v1.AdminQueryReq) (v1.AdminQueryRes, error) {
	return service.Admin.AdminQuery(ctx, req.AdminQuery)
}
