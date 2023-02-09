package constants

/*
* @author Ya peng
* @E-mail linuxsan@msn.com
* @date 2023/1/18 22:31
 */
const (
	CodeSuccess        = 0 // 成功！
	CodeFail           = -1
	CodeAuthticateFail = -2   // 认证失败！
	CodeCreateUserFail = 1001 // 创建用户失败！
	CodeNoSuchUser     = 1002 // 当前用户不存在

	CodeCreateProviderFail = 1010 // 添加IDC公司失败// ！
	CodeGetProviderFail    = 1011 // 无法获取IDC公司列表！
	CodeDelProviderFail    = 1012 // 删除IDC公司失败
	CodeUpdateProviderFail = 1013

	CodeCreateIdcFail = 1010 // 添加机房失败// ！
	CodeGetIdcFail    = 1011 // 无法获取机房列表！
	CodeDelIdcFail    = 1012 // 删除IDC机房失败
	CodeUpdateIdcFail = 1013

	//CodeCreateHostFail       = 1012 // 添加主机信息失败！
	//CodeGetHostFail          = 1013 // 无法获取主机列表
	//CodeDelHostFail          = 1015 // 删除主机失败

	//CodeCreateHostCategoryFail = 2010 // 创建主机类别失败！
	//CodeGetHostCategoryFail    = 2011 // 无法获取主机类别列表！
	//CodeCreateHostFail         = 2012 // 添加主机信息失败！
	//CodeHostCategoryNotExist   = 2013 // 主机类别不存在！
	//CodeGetHostFail            = 2014 // 无法获取主机列表
	//
	//CodeDelHostFail = 1015 // 删除主机失败
)
