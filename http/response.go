package http

// Response 通用HTTP返回值
// Code 0 代表成功，非0代表失败
// ErrMsg 是具体失败信息
// Data 是额外传输信息
type Response struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

const (
	CodeCommonError = 1
)
