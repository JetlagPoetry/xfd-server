package utils

// StringInSlice  判断一个字符串是否在一个数组中
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
