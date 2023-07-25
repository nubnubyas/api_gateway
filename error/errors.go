// Description: 错误定义
package errors

import (
	"github.com/cloudwego/api_gateway/error/kitex_gen/common"
)

// Err is the error type defined in IDL.
type Err struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// New creates a new error with the specified error code.
// The error code must be defined in common
func New(errCode common.Err) Err {
	return Err{
		ErrCode: int64(errCode),
		ErrMsg:  errCode.String(),
	}
}

// Error returns the error message.

func (e Err) Error() string {
	return e.ErrMsg
}
