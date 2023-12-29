package code

const (
	Success       int = 1001
	NoUser        int = 1002
	ServerBusy    int = 1003
	WrongPassWord int = 1004
	InvalidToken  int = 1005
	ExpiredToken  int = 1006
	WithoutToken  int = 1007
	GetError      int = 1008
	OperateFail   int = 1009
	BindError     int = 1010
	ClickBusy     int = 1011
	OutOfMax      int = 1012
	TypeError     int = 1013
)

var msgFlags = map[int]string{
	Success:       "操作成功",
	NoUser:        "用户不存在",
	ServerBusy:    "服务器繁忙",
	WrongPassWord: "用户名或密码错误",
	WithoutToken:  "未携带Token",
	InvalidToken:  "无效的Token",
	ExpiredToken:  "过期的Token",
	GetError:      "获取信息失败",
	OperateFail:   "操作失败",
	BindError:     "JSON绑定失败",
	ClickBusy:     "点击过于频繁",
	OutOfMax:      "上传图片过大",
	TypeError:     "不支持该类型",
}

func Msg(flag int) string {
	return msgFlags[flag]
}
