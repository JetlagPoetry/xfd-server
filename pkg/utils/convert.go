package utils

func IntToBool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}
func IntPtr(i int) *int {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(i bool) *bool {
	return &i
}
