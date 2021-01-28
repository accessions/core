package err

import (
	"fmt"
)

type CodeError struct {
	errCode int
	errMsg string
}

func (c *CodeError) GetErrCode () int {
	return c.errCode
}

func (c *CodeError) GetErrMsg () string {
	return c.errMsg
}

func (c *CodeError) Error () string {
	return fmt.Sprintf("ErrCode:%d, ErrMsg:%s", c.errCode, c.errMsg)
}

func New(errCode int, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg:  errMsg}
}

func NewErrCode(errCode int) *CodeError {
	return &CodeError{errCode: errCode, errMsg: ErrMaps(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: BAD_REQUEST_ERROR, errMsg: errMsg}
}
