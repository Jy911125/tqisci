package common

// 常量定义
const (
	// HTTP 状态码
	HTTPStatusOK                  = 200
	HTTPStatusBadRequest          = 400
	HTTPStatusUnauthorized        = 401
	HTTPStatusInternalServerError = 500

	// 错误信息
	ErrInternalServerError = "Internal Server Error"
	ErrBadRequest          = "Bad Request"
	ErrUnauthorized        = "Unauthorized"
)
