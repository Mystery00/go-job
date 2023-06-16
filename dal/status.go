package dal

const (
	WAIT = 1 + iota
	SCHEDULED
	EXECUTING
	SUCCESS
	FAILED
)

func Status(code int8) string {
	switch code {
	case WAIT:
		return "待执行"
	case SCHEDULED:
		return "已调度"
	case EXECUTING:
		return "执行中"
	case SUCCESS:
		return "执行成功"
	case FAILED:
		return "执行失败"
	default:
		return "未知状态"
	}
}
