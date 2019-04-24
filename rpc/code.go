package main


const (
	RespStatusOK               = 0 // success
	RespStatusArgs             = 4001 // 参数错误
)

// Sms status type
const (
	SmsStatusInit             = "INIT" // Sms status is init.
	SmsStatusVerified         = "VERIFIED" // Sms status is verified.
)

// 短信类型
const (
	SmsTypeInfo                = "INFO" // 公共短信类型.
	SmsTypeReg                 = "REG" // 注册短信类型.
	SmsTypeLogin               = "LOGIN" // 登录短信类型.
	SmsTypeForget              = "FORGET" // 忘记密码，重置短信。
)
