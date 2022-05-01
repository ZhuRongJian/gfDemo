package response

import (
	"github.com/gogf/gf/v2/errors/gerror"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Response(data interface{}, err error) Result {
	if err != nil {
		return Error(err)
	}
	return Success(data, "")
}
func Success(data interface{}, message string) Result {
	return Result{
		Code:    0,
		Data:    data,
		Message: message,
	}
}
func Error(err error) Result {
	code := gerror.Code(err)
	return Result{
		Code:    code.Code(),
		Message: err.Error(),
	}
}
