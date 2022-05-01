package v1

import "gfDemo/utility/query"

type BaseEmptyRes struct {
}

type BasePageQueryRes struct {
	*query.Result
}