package http

import (
	"encoding/json"
	"net/http"
)

type AppErr struct {
	HttpCode int
	Code     string
	Msg      map[string]string
}

// i18n support
var (
	ENGLISH = "en_US"
	CHINESE = "zh_CN"
)

var (
	ErrLoginFailed              = &AppErr{HttpCode: http.StatusOK, Code: "1000", Msg: map[string]string{CHINESE: "用户名或密码错误", ENGLISH: "incorrect username or password"}}
	ErrAuthFailed               = &AppErr{HttpCode: http.StatusUnauthorized, Code: "1001", Msg: map[string]string{CHINESE: "鉴权失败", ENGLISH: "authentication failed"}}
	ErrInvalidParam             = &AppErr{HttpCode: http.StatusBadRequest, Code: "1002", Msg: map[string]string{CHINESE: "请求参数不正确", ENGLISH: "invalid request parameter"}}
	ErrInvalidToken             = &AppErr{HttpCode: http.StatusOK, Code: "1003", Msg: map[string]string{CHINESE: "鉴权失败，请传入token", ENGLISH: "invalid token"}}
	ErrRequestValidationNotPass = &AppErr{HttpCode: http.StatusOK, Code: "1005", Msg: map[string]string{CHINESE: "请求有效性检查未通过", ENGLISH: "request validation not pass"}}
	ErrDataNotExist             = &AppErr{HttpCode: http.StatusOK, Code: "1006", Msg: map[string]string{CHINESE: "记录不存在", ENGLISH: "record not exist"}}
	ErrDataAlreadyExist         = &AppErr{HttpCode: http.StatusOK, Code: "1007", Msg: map[string]string{CHINESE: "记录已存在", ENGLISH: "record already exist"}}
	ErrTryLater                 = &AppErr{HttpCode: http.StatusOK, Code: "1008", Msg: map[string]string{CHINESE: "系统繁忙，请稍后重试", ENGLISH: "system busy, please try again later"}}
	ErrMethodNotAllowed         = &AppErr{HttpCode: http.StatusMethodNotAllowed, Code: "1009", Msg: map[string]string{CHINESE: "不支持的请求方法", ENGLISH: "method not allowed"}}
	ErrLogin                    = &AppErr{HttpCode: http.StatusOK, Code: "2001", Msg: map[string]string{CHINESE: "token失效，请重新登陆", ENGLISH: "token expired, please login again"}}
	ErrAlreadyLogin             = &AppErr{HttpCode: http.StatusOK, Code: "2002", Msg: map[string]string{CHINESE: "请勿重复登录", ENGLISH: "already login in"}}
	ErrEmailAlreadyExist        = &AppErr{HttpCode: http.StatusOK, Code: "2003", Msg: map[string]string{CHINESE: "邮箱已存在", ENGLISH: "email already exist"}}
	ErrUsernameAlreadyExist     = &AppErr{HttpCode: http.StatusOK, Code: "2004", Msg: map[string]string{CHINESE: "用户名已存在", ENGLISH: "username already exist"}}
	ErrPasswordNotTheSame       = &AppErr{HttpCode: http.StatusOK, Code: "2005", Msg: map[string]string{CHINESE: "新密码不一致", ENGLISH: "password not the same"}}
	ErrOldPasswordNotRight      = &AppErr{HttpCode: http.StatusOK, Code: "2006", Msg: map[string]string{CHINESE: "旧密码不正确", ENGLISH: "wrong password"}}
	ErrPhoneNotExited           = &AppErr{HttpCode: http.StatusOK, Code: "2007", Msg: map[string]string{CHINESE: "手机号不存在", ENGLISH: "phone number not exist"}}
	ErrUserNameNotExited        = &AppErr{HttpCode: http.StatusOK, Code: "2008", Msg: map[string]string{CHINESE: "用户名不存在", ENGLISH: "username not exist"}}
	ErrInvalidPassword          = &AppErr{HttpCode: http.StatusOK, Code: "2009", Msg: map[string]string{CHINESE: "密码错误", ENGLISH: "wrong password"}}
	ErrBusinessCheckNotPass     = &AppErr{HttpCode: http.StatusOK, Code: "2010", Msg: map[string]string{CHINESE: "业务检查未通过", ENGLISH: "business validation not pass"}}
	ErrInternal                 = &AppErr{HttpCode: http.StatusInternalServerError, Code: "9999", Msg: map[string]string{CHINESE: "系统错误", ENGLISH: "system error"}}
)

func ResponseOK(w http.ResponseWriter, data interface{}) {
	if data == nil {
		data = ""
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"code": "0000", "msg": "", "data": data})
}

func ResponseErr(w http.ResponseWriter, locale string, appErr *AppErr, reason ...string) {
	w.WriteHeader(appErr.HttpCode)
	w.Header().Set("Content-Type", "application/json")
	msg := getMsg(locale, appErr)
	if len(reason) > 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": appErr.Code, "msg": msg, "reason": reason[0], "data": ""})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"code": appErr.Code, "msg": msg, "data": ""})
	}
}

func getMsg(locale string, err *AppErr) string {
	if locale == "" {
		locale = ENGLISH
	}
	msg, ok := err.Msg[locale]
	if !ok {
		msg, ok = err.Msg[ENGLISH]
		if !ok {
			msg = ""
		}
	}
	return msg
}
