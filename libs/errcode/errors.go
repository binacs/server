package errcode

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
)

// ---------------------------- web ------------------------------

// Response the response
type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// WebResponse the web response
func WebResponse(err *Error, data interface{}) *Response {
	return &Response{
		Code: err.Code(),
		Msg:  err.Error(),
		Data: data,
	}
}

// ---------------------------- grpc ------------------------------
var (
	grpcErrorMake  *ErrorMaker = NewErrorMaker(2002)
	ErrGrpcSuccess             = grpcErrorMake.MakeSuccess()

	ErrGrpcAuth = grpcErrorMake.MakeAppError(int(codes.Unauthenticated), "鉴权失败", "auth error")

	ErrGrpcAppExecute         = grpcErrorMake.MakeAppError(101, "执行出错", "Execute error")
	ErrGrpcAppAccTokenExpired = grpcErrorMake.MakeAppError(102, "TOKEN过期", "AccessToken expired")

	ErrGrpcSysRedisErr = grpcErrorMake.MakeSystemError(1, "Redis错误", "redis error")
	ErrGrpcSysMysqlErr = grpcErrorMake.MakeSystemError(2, "Mysql错误", "mysql error")
)

// ---------------------------- Error ------------------------------

// Error error
type Error struct {
	code        int64
	chineseDesc string
	englishDesc string
	detail      string
}

// Code return the code
func (e *Error) Code() int64 {
	return e.code
}

// Msg make the msg
func (e *Error) Msg(lang string) string {
	switch strings.ToUpper(lang) {
	case LangChinese:
		if len(e.detail) == 0 {
			return e.chineseDesc
		}
		return e.chineseDesc + "[" + e.detail + "]"
	default:
		if len(e.detail) == 0 {
			return e.englishDesc
		}
		return e.englishDesc + "[" + e.detail + "]"
	}
}

// Desc return the desc
func (e *Error) Desc(lang string) string {
	switch strings.ToUpper(lang) {
	case LangChinese:
		return e.chineseDesc
	default:
		return e.englishDesc
	}
}

// Error return the string(msg)
func (e *Error) Error() string {
	return e.Msg(LangChinese)
}

// AppendMsg append msg
func (e *Error) AppendMsg(f interface{}, v ...interface{}) *Error {
	if !debugMode {
		return e
	}
	return &Error{e.code, e.chineseDesc, e.englishDesc, Format(f, v...)}
}

// ------------------------ ErrorMaker--------------------------
var (
	LangEnglish = "EN"
	LangChinese = "CN"
)

var debugMode = true

// SetDebugMode set debug mode
func SetDebugMode(mode bool) {
	debugMode = mode
}

// ErrorMaker error maker
type ErrorMaker struct {
	serverid   int
	appErrList map[int64]interface{}
	sysErrList map[int64]interface{}
}

// NewErrorMaker return a pointer to ErrorMaker
func NewErrorMaker(serverid int) *ErrorMaker {
	return &ErrorMaker{
		serverid:   serverid,
		appErrList: make(map[int64]interface{}),
		sysErrList: make(map[int64]interface{}),
	}
}

// SetServerid set the serverid
func (m *ErrorMaker) SetServerid(serverid int) bool {
	m.serverid = serverid
	return true
}

// GetServiceID get the serverid
func (m *ErrorMaker) GetServerID() int {
	return m.serverid
}

// GetAppErrList get the app error list
func (m *ErrorMaker) GetAppErrList() map[int64]interface{} {
	return m.appErrList
}

// GetSysErrList get the system error list
func (m *ErrorMaker) GetSysErrList() map[int64]interface{} {
	return m.sysErrList
}

// MakeSystemError make system error
func (m *ErrorMaker) MakeSystemError(errorseq int, ChineseMsg string, EnglishMsg string) *Error {
	const typeSystem = 1
	errcode := int64(m.serverid*10000 + typeSystem*1000 + errorseq)
	err := &Error{errcode, ChineseMsg, EnglishMsg, ""}
	errinfo := make(map[string]string)
	errinfo["chineseDesc"] = ChineseMsg
	errinfo["englishDesc"] = EnglishMsg
	m.sysErrList[errcode] = errinfo
	return err
}

// MakeAppError make app error
func (m *ErrorMaker) MakeAppError(errorseq int, ChineseDesc string, EnglishDesc string) *Error {
	const typeSystem = 2
	errcode := int64(m.serverid*10000 + typeSystem*1000 + errorseq)
	err := &Error{errcode, ChineseDesc, EnglishDesc, ""}
	errinfo := make(map[string]string)
	errinfo["chineseDesc"] = ChineseDesc
	errinfo["englishDesc"] = EnglishDesc
	m.appErrList[errcode] = errinfo
	return err
}

// MakeSuccess make success
func (m *ErrorMaker) MakeSuccess() *Error {
	return &Error{0, "成功", "Success", ""}
}

// Format common usage
func Format(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format
		} else {
			//not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
