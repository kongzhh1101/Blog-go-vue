package res

type ErrorCode int

const (
	SettingError   ErrorCode = 1001 // 系统错误
	ArguementError ErrorCode = 1002 // 参数错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingError:   "系统错误",
		ArguementError: "参数错误",
	}
)
