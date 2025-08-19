package models

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 成功响应
func SuccessResponse(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// 成功响应带自定义消息
func SuccessWithMessage(message string, data interface{}) Response {
	return Response{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

// 错误响应
func ErrorResponseWithCode(code int, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// 错误响应带数据
func ErrorResponseWithData(code int, message string, data interface{}) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// 常用错误响应
func BadRequestResponse(message string) ErrorResponse {
	return ErrorResponseWithCode(400, message)
}

func UnauthorizedResponse(message string) ErrorResponse {
	return ErrorResponseWithCode(401, message)
}

func ForbiddenResponse(message string) ErrorResponse {
	return ErrorResponseWithCode(403, message)
}

func NotFoundResponse(message string) ErrorResponse {
	return ErrorResponseWithCode(404, message)
}

func InternalServerErrorResponse(message string) ErrorResponse {
	return ErrorResponseWithCode(500, message)
}