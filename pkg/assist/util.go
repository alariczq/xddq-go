package assist

import (
	"github.com/golang-module/carbon/v2"
)

func IsToday(t int64) bool {
	return carbon.CreateFromTimestampMilli(t).IsToday()
}

func TomorrowMilli() int64 {
	return carbon.Now().AddDays(1).StartOfDay().TimestampMilli()
}
