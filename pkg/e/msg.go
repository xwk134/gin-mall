package e

var MsgFlags = map[int]string{
	Success:               "ok",
	Error:                 "fail",
	InvalidParams:         "参数错误",
	ErrorExisUser:         "用户名已存在",
	ErrorFailEncryption:   "密码加密失败",
	ErrorExisUserNotFound: "用户不存在",
	ErrorNotCompare:       "密码错误",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
