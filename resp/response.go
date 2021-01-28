package resp

import "github.com/accessions/core/resp/err"

type (
	NullJson struct {}

	ResponseSuccessBean struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Data interface{} `json:"data"`
	}

	ResponseFailureBean struct {
		Code int `json:"code"`
		Msg string `json:"message"`
	}
)

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{err.OK, "OK", data}
}

func Failure(errCode int) *ResponseFailureBean {
	code := err.NewErrCode(errCode)
	return &ResponseFailureBean{code.GetErrCode(), code.GetErrMsg()}
}
