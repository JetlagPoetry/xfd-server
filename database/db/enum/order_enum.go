package enum

type ModifyQuantityType int

const (
	ModifyQuantityAdd    ModifyQuantityType = 1
	ModifyQuantityReduce ModifyQuantityType = 2
)

func (m ModifyQuantityType) Code() int {
	switch m {
	case ModifyQuantityAdd:
		return 1
	case ModifyQuantityReduce:
		return 2
	default:
		return 0
	}
}

func (m ModifyQuantityType) Info() (int, string) {
	switch m {
	case ModifyQuantityAdd:
		return 1, "增加数量"
	case ModifyQuantityReduce:
		return 2, "减少数量"
	default:
		return 0, "错误"
	}
}

func (m ModifyQuantityType) IsValid() bool {
	return m == ModifyQuantityAdd || m == ModifyQuantityReduce
}
