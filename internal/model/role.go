package model

import (
	"gfDemo/utility/query"
	"github.com/gogf/gf/v2/database/gdb"
	"strings"
)

type RoleQuery struct {
	query.PageParam
	PartnerType string `json:"partner_type"`
	IsAdmin     uint   `json:"is_admin"`
	Status      uint   `json:"status"`
}

func (q *RoleQuery) Cond(m *gdb.Model) *gdb.Model {
	if q.PartnerType != "" {
		m = m.Where("partner_type =?", strings.TrimSpace(q.PartnerType))
	}
	if q.IsAdmin > 0 {
		m = m.Where("is_main =?", q.IsAdmin)
	}
	if q.Status > 0 {
		m = m.Where("status =?", q.Status)
	}
	return m
}

type RoleListQuery struct {
	query.PageParam
}

func (q *RoleListQuery) Cond(m *gdb.Model) *gdb.Model {
	return m
}

type RoleVo struct {
	Id        uint   `json:"id"`
	Name      string `json:"name" v:"required#角色名称不能空"`
	Code      string `json:"code"`
	Sort      int    `json:"sort"`
	Note      string `json:"note"`
	UpdatedBy uint64
	CreatedBy uint64
}
