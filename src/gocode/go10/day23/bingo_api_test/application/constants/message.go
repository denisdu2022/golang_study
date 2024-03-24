package constants

const (
	//用户相关

	Success           = "成功!"
	CreateUserSuccess = "创建用户成功!"
	CreateUserFail    = "创建用户失败!"
	NoSuchUser        = "用户不存在!"
	PasswordError     = "密码错误!"

	//Token相关

	TokenIsExpired       = "认证令牌过期!"
	TokenIsNotValidYet   = "令牌尚未激活!"
	TokenIsMalformed     = "认证令牌格式有误!"
	TokenIsInvalid       = "无效的认证令牌!"
	HostCategoryNotExist = "主机类别不存在！"

	AddPublicKeyFail = "添加公钥失败!"

	HostInfoError   = "主机信息有误！！"
	SSHKeyIsInvalid = "无效的秘钥信息！"
	MissingKeys     = "丢失公钥或者私钥"

	ParamErrorMsg     = "客户端请求参数有误!"
	ExecuteCMDFailMsg = "执行命令失败!"
)
