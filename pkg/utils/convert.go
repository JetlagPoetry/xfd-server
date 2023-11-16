package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
)

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

func Float64Ptr(i float64) *float64 {
	return &i
}

func Float32Ptr(i float32) *float32 {
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

func StringToDecimal(i string) decimal.Decimal {
	if i == "" {
		return decimal.NewFromFloat(0.0)
	}
	dec, err := decimal.NewFromString(i)
	if err != nil {
		fmt.Println("无法解析字符串为 decimal.Decimal:", err)
		return decimal.NewFromFloat(0.0)
	}

	return dec
}
