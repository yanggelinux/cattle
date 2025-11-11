package ce

import (
	"github.com/yanggelinux/cattle/common/coder"
	"net/http"
)

func NewCoder(code, httpStatus int, msg string) coder.Coder {
	return coder.DefaultCoder{C: code, HTTP: httpStatus, Msg: msg}
}

var (
	Success            = NewCoder(200, http.StatusOK, "请求成功")
	Error              = NewCoder(500, http.StatusOK, "请求错误")
	ErrorInvaildParams = NewCoder(4001, http.StatusOK, "请求参数错误")

	ErrorAuthCheckTokenFailed  = NewCoder(10001, http.StatusOK, "Token鉴权失败")
	ErrorAuthCheckTokenTimeout = NewCoder(10002, http.StatusOK, "Token已超时")
	ErrorAuthTokenFailed       = NewCoder(10003, http.StatusOK, "请求Token缺失")
	ErrorAuthToken             = NewCoder(10004, http.StatusOK, "Token生成失败")
	ErrorAuthCheckFailed       = NewCoder(10005, http.StatusOK, "鉴权失败")
	ErrorGetToken              = NewCoder(10006, http.StatusOK, "获取Token失败")
	ErrorGetAuthor             = NewCoder(10007, http.StatusOK, "获取Authorization失败")

	ErrorMetaInfoFailed = NewCoder(11001, http.StatusOK, "请求元数据错误")
	ErrorAPICheckFailed = NewCoder(12001, http.StatusOK, "API权限验证失败")

	ErrorLoginUserPasswordEmpty = NewCoder(21001, http.StatusOK, "用户名或密码为空")
	ErrorLoginUserFailed        = NewCoder(21002, http.StatusOK, "用户名或密码错误")
	ErrorLoginFailed            = NewCoder(21004, http.StatusOK, "登录失败")
	ErrorLoginLDAPFailed        = NewCoder(21004, http.StatusOK, "LDAP认证失败")

	ErrorQueryFailed   = NewCoder(30001, http.StatusOK, "查询失败")
	ErrorQueryNotFound = NewCoder(30002, http.StatusOK, "查询记录不存在")

	ErrorDBQueryFailed       = NewCoder(40002, http.StatusOK, "数据库查询错误")
	ErrorDBOperateFailed     = NewCoder(40003, http.StatusOK, "数据库操作错误")
	ErrorDBQueryNotFound     = NewCoder(40004, http.StatusOK, "数据库记录不存在")
	ErrorDBDuplicateEntry    = NewCoder(40005, http.StatusOK, "数据重复，已经存在")
	ErrorOrderDuplicateLabel = NewCoder(40020, http.StatusOK, "工单标签重复，已经存在")

	ErrorValidateForm      = NewCoder(50001, http.StatusOK, "表单数据校验失败")
	ErrorUploadFile        = NewCoder(50010, http.StatusOK, "上传文件失败")
	ErrorUploadExceedLimit = NewCoder(50011, http.StatusOK, "上传文件失败,超出限制")

	ErrorGroupHasRecord      = NewCoder(60001, http.StatusOK, "架构组中包含未删除组")
	ErrorGroupHasGraphRecord = NewCoder(60002, http.StatusOK, "架构组中包含未删除图")
	ErrorRelGraphNode        = NewCoder(60003, http.StatusOK, "关联架构图包含多个开始节点或结束节点")

	ErrorDecryptPassword = NewCoder(70001, http.StatusOK, "密码解密失败")
	ErrorParsePassword   = NewCoder(70001, http.StatusOK, "密码解析失败")

	ErrorRequestSeal = NewCoder(80001, http.StatusOK, "调用Seal接口失败")
)

func registerCoder() {
	coder.Register(Success)
	coder.Register(Error)
	coder.Register(ErrorInvaildParams)
	coder.Register(ErrorAuthCheckTokenFailed)
	coder.Register(ErrorAuthCheckTokenTimeout)
	coder.Register(ErrorAuthTokenFailed)
	coder.Register(ErrorAuthToken)
	coder.Register(ErrorAuthCheckFailed)

	coder.Register(ErrorGetToken)
	coder.Register(ErrorGetAuthor)

	coder.Register(ErrorMetaInfoFailed)
	coder.Register(ErrorAPICheckFailed)
	coder.Register(ErrorLoginUserPasswordEmpty)
	coder.Register(ErrorLoginUserFailed)
	coder.Register(ErrorLoginFailed)
	coder.Register(ErrorLoginLDAPFailed)

	coder.Register(ErrorDBQueryFailed)
	coder.Register(ErrorDBOperateFailed)
	coder.Register(ErrorDBQueryNotFound)
	coder.Register(ErrorQueryFailed)
	coder.Register(ErrorQueryNotFound)
	coder.Register(ErrorDBDuplicateEntry)
	coder.Register(ErrorOrderDuplicateLabel)

	coder.Register(ErrorValidateForm)
	coder.Register(ErrorUploadFile)
	coder.Register(ErrorUploadExceedLimit)

	coder.Register(ErrorGroupHasRecord)
	coder.Register(ErrorGroupHasGraphRecord)
	coder.Register(ErrorRelGraphNode)
	coder.Register(ErrorDecryptPassword)
	coder.Register(ErrorParsePassword)
	coder.Register(ErrorRequestSeal)

}

// 包内初始化加载
func init() {
	registerCoder()
}
