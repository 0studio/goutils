package goutils

import (
	"time"
)

const (
	DAY_SECONDS    = 24 * 60 * 60
	HOUR_SECONDS   = 60 * 60
	MINUTE_SECONDS = 60
)

var UninitedTime time.Time // 主要用来比较 某个时间 是 否初始化过
