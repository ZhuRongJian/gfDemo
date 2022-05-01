package model

import (
	"gfDemo/utility/query"
	"github.com/gogf/gf/v2/database/gdb"
	"strings"
)

type ConfigQuery struct {
	query.PageParam
}

func (q *ConfigQuery) Cond(m *gdb.Model) *gdb.Model {
	q.Order = "asc"
	q.Sort = "sort"
	return m
}

type ConfigVo struct {
	Id        uint64   `json:"id"`                        // 主键ID
	Name      string `json:"name" v:"required#配置名称不能空"` // 配置名称
	Code      string `json:"code" v:"required#配置编码不能空"` // 配置编码
	Sort      int    `json:"sort" v:"required#排序号不能空"`  // 排序
	Note      string `json:"note"`                      // 备注
	CreatedBy uint64   `json:"created_by"`                // 添加人
	UpdatedBy uint64   `json:"updated_by"`                // 更新人
}

type ConfigDataQuery struct {
	query.PageParam
	ConfigId uint   `json:"config_id"`
	Title    string `json:"title"`
	Code     string `json:"code"`
}

func (q *ConfigDataQuery) Cond(m *gdb.Model) *gdb.Model {
	q.Order = "asc"
	q.Sort = "sort"
	if q.ConfigId > 0 {
		m = m.Where("config_id=?", q.ConfigId)
	}
	if q.Title != "" {
		m = m.WhereLike("title", "%"+strings.TrimSpace(q.Title)+"%")
	}
	if q.Code != "" {
		m = m.WhereLike("code", "%"+strings.TrimSpace(q.Code)+"%")
	}
	return m
}

type ConfigDataVo struct {
	Id        uint64   `json:"id"         `                // 主键ID
	ConfigId  uint64   `json:"config_id"  `                // 配置ID
	Title     string `json:"title" v:"required#配置标题不能空"` // 配置标题
	Code      string `json:"code" v:"required#配置编码不能空"`  // 配置编码
	Value     string `json:"value"      `                // 配置值
	Status    uint64   `json:"status"     `                // 状态：1正常 2停用
	Sort      uint64   `json:"sort"       `                // 排序
	Note      string `json:"note"       `                // 配置说明
	CreatedBy uint64   `json:"created_by" `                // 添加人
	UpdatedBy uint64   `json:"updated_by" `                // 更新人
}
