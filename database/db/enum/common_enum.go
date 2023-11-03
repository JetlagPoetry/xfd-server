package enum

type RequestSourceType int32

const (
	MiniProgramSource RequestSourceType = 1
	BackendSource     RequestSourceType = 2
)

func (r RequestSourceType) Info() (int, string) {
	switch r {
	case MiniProgramSource:
		return 1, "小程序"
	case BackendSource:
		return 2, "后台"
	default:
		return 0, "error"
	}
}

func (r RequestSourceType) Code() int {
	switch r {
	case MiniProgramSource:
		return 1
	case BackendSource:
		return 2
	default:
		return 0
	}
}
