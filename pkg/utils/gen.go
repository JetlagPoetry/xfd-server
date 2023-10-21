package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func GenSixDigitCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func GenUUID() string {
	return uuid.New().String()
}
