package err

var message map[int]string

func init()  {
	message = make(map[int]string)
	message[OK] = "SUCCESS"
	message[BAD_REQUEST_ERROR] = "服务器繁忙,请稍后再试"
	message[REQUEST_PARAM_ERROR] = "参数错误"
	message[USER_NOT_FOUND] = "用户不存在"
}

func ErrMaps(errcode int) string  {
	if msg, ok := message[errcode]; ok {
		return msg
	}
	return "服务器繁忙,请稍后再试"
}
