package constants

const (
	CodeSuccess        = 0    //成功
	CodeFail           = -1   //失败
	CodeAuthticateFail = -2   //用户认证失败
	CodeCreateUserFail = 1001 //创建用户失败
	CodeNoSuchUser     = 1002 //当前用户不存在

	CodeCreateHostCategoryFail = 1010 // 创建主机类别失败！
	CodeGetHostCategoryFail    = 1011 // 无法获取主机类别列表！
	CodeCreateHostFail         = 1012 // 添加主机信息失败！
	CodeHostCategoryNotExist   = 1013 // 主机类别不存在！
	CodeGetHostFail            = 1014 // 无法获取主机列表
	CodeDelHostFail            = 1015 //删除主机失败

	ParamError     = -3   //客户端请求参数有误
	ExecuteCMDFail = 1016 //执行命令失败

	CodeCreateCmdCategoryFail = 1020 //创建指令模板类别失败
	CodeGetCmdCategoryFail    = 1021 //无法获取指令模板类别列表!
	CodeCreateCmdFail         = 1022 //添加指令模板信息失败
	CodeGetCmdFail            = 1023 //指令模板不存在
)
