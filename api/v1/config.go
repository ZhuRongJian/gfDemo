package v1

import (
	"gfDemo/internal/model"
	"gfDemo/utility/query"
	"github.com/gogf/gf/v2/frame/g"
)

type ConfigQueryReq struct {
	g.Meta `path:"/config/query" tags:"Config" method:"get" summary:"获取配置值列表"`
	*model.ConfigQuery
}

type ConfigQueryRes struct {
	*query.Result
}

type ConfigChangeReq struct {
	g.Meta `path:"/config/change" tags:"Config" method:"post" summary:"修改/新增配置值"`
	*model.ConfigVo
}

type ConfigDeleteReq struct {
	g.Meta `path:"/config/delete" tags:"Config" method:"post" summary:"删除配置值"`
	Id     uint `json:"id"`
}

type ConfigDataQueryReq struct {
	g.Meta `path:"/config_data/query" tags:"ConfigData" method:"get" summary:"获取配置类型列表"`
	*model.ConfigDataQuery
}


type ConfigDataChangeReq struct {
	g.Meta `path:"/config_data/change" tags:"ConfigData" method:"post" summary:"修改/新增配置类型"`
	*model.ConfigDataVo
}

type ConfigDataDeleteReq struct {
	g.Meta `path:"/config_data/delete" tags:"ConfigData" method:"post" summary:"删除配置类型"`
	Id     uint `json:"id"`
}
