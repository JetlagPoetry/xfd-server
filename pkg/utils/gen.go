package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

func GenSixDigitCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func GenUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
func GenUsername(phone string) string {
	return "用户" + phone
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

func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

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

func TimeFormatUs() string {
	// 当前时间对象
	curTime := time.Now()
	// curTime的时间戳（秒）
	unixS := curTime.Unix()
	// curTime的时间戳（毫秒）
	unixMs := curTime.UnixNano() / 1e6
	// curTime的时间戳（微秒）
	unixUs := curTime.UnixNano() / 1e3
	// 毫秒时间
	timeMs := unixMs - unixS*1e3
	// 如果毫秒数不够三位的话，则在前面补0
	msStr := sup(timeMs, 3)
	// 微妙时间
	timeUs := unixUs - unixMs*1e3
	// 如果微秒数不够三位的话，则在前面补0
	usStr := sup(timeUs, 3)
	// curTime的日期格式
	dateStr := curTime.Format("20060102150405")

	return dateStr + msStr + usStr
}

func sup(i int64, n int) string {
	msStr := strconv.FormatInt(i, 10)
	for len(msStr) < n {
		msStr = "0" + msStr
	}
	return msStr
}

var num int64

func GenerateOrder() string {
	t := time.Now().Format("20060102150405")
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s", t, ps, rs)
	if num > 9999999999 {
		num = 0
	}
	return n
}

func GenerateRandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}
