package errors

const (
	// CodeSuccess 成功
	CodeSuccess = 200
	// CodeBadRequest 通用参数错误
	CodeBadRequest = 400
	// CodeUnauthorized 未认证
	CodeUnauthorized = 401
	// CodeForbidden 无权限访问
	CodeForbidden = 403
	// CodeNotFound 资源不存在
	CodeNotFound = 404
	// CodeServerError 服务端错误
	CodeServerError = 500

	// ========== JWT模块 (10000 - 10100) ==========
	CodeAccessTokenExpired  = 10001
	CodeAccessTokenInvalid  = 10002
	CodeRefreshTokenInvalid = 10003
	CodeRefreshTokenExpired = 10004

	// ========== 用户模块 (2101-2199) ==========
	CodeUserNotFound          = 2101
	CodeUsernameOrPasswordErr = 2102
	CodeUsernameAlreadyExists = 2103
	CodeUserStatusNotValid    = 2104
	CodeUserDisabled          = 2105

	// ========== 文章模块 (2201-2299) ==========
	CodeArticleNotFound = 2201

	// ========== 评论模块 (2301-2399) ==========
	CodeCommentNotFound = 2301

	// ========== 分类模块 (2401-2499) ==========
	CodeCategoryNotFound = 2401

	// ========== 标签模块 (2501-2599) ==========
	CodeTagNotFound          = 2501
	CodeTagNameAlreadyExists = 2502

	// ========== 验证码模块 (2601-2699) ==========
	CodeCaptchaError    = 2601
	CodeCaptchaExpired  = 2602
	CodeCaptchaRequired = 2603
)

// codeMessages 错误码对应的默认消息
var codeMessages = map[int]string{
	CodeSuccess:               "成功",
	CodeBadRequest:            "请求参数错误",
	CodeUnauthorized:          "未认证",
	CodeForbidden:             "无权限访问",
	CodeNotFound:              "资源不存在",
	CodeServerError:           "服务器内部错误",
	CodeAccessTokenExpired:    "访问令牌已过期",
	CodeAccessTokenInvalid:    "访问令牌无效",
	CodeRefreshTokenInvalid:   "刷新令牌无效",
	CodeRefreshTokenExpired:   "刷新令牌已过期",
	CodeUserNotFound:          "用户不存在",
	CodeUsernameOrPasswordErr: "用户名或密码错误",
	CodeUsernameAlreadyExists: "用户名已存在",
	CodeUserStatusNotValid:    "用户状态不合法",
	CodeUserDisabled:          "用户已被禁用",
	CodeArticleNotFound:       "文章不存在",
	CodeCommentNotFound:       "评论不存在",
	CodeCategoryNotFound:      "分类不存在",
	CodeTagNotFound:           "标签不存在",
	CodeTagNameAlreadyExists:  "标签名已存在",
	CodeCaptchaError:          "验证码错误",
	CodeCaptchaExpired:        "验证码已过期",
	CodeCaptchaRequired:       "请获取验证码",
}

// GetMessage 根据错误码获取默认消息
func GetMessage(code int) string {
	if msg, ok := codeMessages[code]; ok {
		return msg
	}
	return "未知错误"
}
