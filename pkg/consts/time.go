package consts

import "time"

var TimeZeroValue, _ = time.ParseInLocation("2006-01-02 15:04:05", "1000-01-01 00:00:00", time.Local)
