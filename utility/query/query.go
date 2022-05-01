package query

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"math"
)

//Query 分页查询条件
type Query interface {
	GetPageIndex() int
	GetPageSize() int
	Cond(*gdb.Model) *gdb.Model
	GetOrder() string
}

//PageParam 分页参数
type PageParam struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

//GetPageIndex 获取当前页码
func (p *PageParam) GetPageIndex() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

//GetPageSize 获取分页大小
func (p *PageParam) GetPageSize() int {
	if p.Limit == 0 {
		return 20
	}
	if p.Limit > 100 {
		return 100
	}
	return p.Limit
}

//Asc 正序排
func (p *PageParam) Asc(column string) {
	p.Sort = column
	p.Order = "asc"
}

//Desc 倒序排
func (p *PageParam) Desc(column string) {
	p.Sort = column
	p.Order = "desc"
}

//GetOrder 获取排序
func (p *PageParam) GetOrder() string {
	if len(p.Sort) > 0 && len(p.Order) > 0 {
		return fmt.Sprintf("%s %s", p.Sort, p.Order)
	}

	return ""
}

//Result 分页结果
type Result struct {
	Records interface{} `json:"records"`
	Current int         `json:"current"`
	Pages   int         `json:"pages"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
}

func (r *Result) WithRecords(data interface{}) *Result {
	r.Records = data
	return r
}
func NewResult(page, limit, total int) *Result {
	pages := int(math.Ceil(float64(total) / float64(limit)))
	return &Result{
		Records: nil,
		Current: page,
		Pages:   pages,
		Size:    limit,
		Total:   total,
	}
}

//Page 分页查询
func Page(m *gdb.Model, query Query, bean interface{}, fields ...string) (*Result, error) {
	m = query.Cond(m)
	mp := m.Clone()
	var total int
	mp.FieldCount("id")
	total, err := mp.Fields("*").Count()
	if err != nil {
		return nil, err
	}

	pageIndex := query.GetPageIndex()
	pageSize := query.GetPageSize()

	for _, v := range fields {
		m = m.Fields(v)
	}
	listM := m.Page(query.GetPageIndex(), query.GetPageSize())
	order := query.GetOrder()
	if len(order) > 0 {
		listM = listM.Order(order)
	}

	err = listM.Scan(bean)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeDbOperationError, err.Error())
	}
	return NewResult(pageIndex, pageSize, total).WithRecords(bean), nil
}

//Exec 执行
func Exec(m *gdb.Model, query Query, bean interface{}, fields ...string) error {
	m = query.Cond(m)
	for _, v := range fields {
		m = m.Fields(v)
	}
	order := query.GetOrder()
	if len(order) > 0 {
		m = m.Order(order)
	}
	err := m.Scan(bean)
	if err != nil {
		return err
	}
	return nil
}
