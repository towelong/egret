package errors

import "net/http"

var (
	UnknownCode      int64 = 9999
	OKCode           int64 = 0
	NotFoundCode     int64 = 1
	UnauthorizedCode int64 = 2
	ForbiddenCode    int64 = 3
	ParamsCode       int64 = 4
)

var codeMessage = map[int64]string{
	UnknownCode:      "服务器未知异常",
	OKCode:           "操作成功",
	NotFoundCode:     "资源未找到",
	UnauthorizedCode: "未授权无法访问",
	ForbiddenCode:    "禁止访问",
	ParamsCode:       "参数异常",
}

func Code2Message(code int64) string {
	return codeMessage[code]
}

var (
	Unknown      = New(UnknownCode, http.StatusInternalServerError, Code2Message(UnknownCode))
	OK           = New(OKCode, http.StatusOK, Code2Message(OKCode))
	NotFound     = New(NotFoundCode, http.StatusNotFound, Code2Message(NotFoundCode))
	Unauthorized = New(UnauthorizedCode, http.StatusUnauthorized, Code2Message(UnauthorizedCode))
	Forbidden    = New(ForbiddenCode, http.StatusForbidden, Code2Message(ForbiddenCode))
	ParamsError  = New(ParamsCode, http.StatusBadRequest, Code2Message(ParamsCode))
)
