package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	// code= 1000... 用户模塊的錯誤
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code= 2000... 文章模塊的錯誤

	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... 分類模塊的錯誤
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密碼錯誤",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,請重新登入",
	ERROR_TOKEN_RUNTIME:    "TOKEN已過期,請重新登入",
	ERROR_TOKEN_WRONG:      "TOKEN不正確,請重新登入",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式錯誤,請重新登入",
	ERROR_USER_NO_RIGHT:    "該用戶無權限",

	ERROR_ART_NOT_EXIST: "文章不存在",
	ERROR_CATENAME_USED:  "該分類已存在",
	ERROR_CATE_NOT_EXIST: "該分類不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
