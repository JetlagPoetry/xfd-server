package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

func GenSixDigitCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func GenUUID() string {
	return uuid.New().String()
}

func GenerateFileName() string {
	currentTime := time.Now()
	yearMonthDay := currentTime.Format("20060102")
	hourMinute := currentTime.Format("1504")
	filename := fmt.Sprintf("%s/%s/%s/%s", yearMonthDay, hourMinute, GenerateRandCode("", 4), GenerateRandCode("", 8))
	return filename
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// GenerateSPU 生成指定长度的SKU编码
func GenerateRandCode(prefix string, length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < 4; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return prefix + string(b) + sb.String()
}
