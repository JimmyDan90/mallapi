package common

import "time"

// 时间格式化
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func WeekTime(days int) (dayTime string) {
	switch days {
	case 0:
		dayTime = BeforeTime(0 * time.Hour)
	case 1:
		dayTime = BeforeTime(-24 * time.Hour)
	case 2:
		dayTime = BeforeTime(-48 * time.Hour)
	case 3:
		dayTime = BeforeTime(-72 * time.Hour)
	case 4:
		dayTime = BeforeTime(-96 * time.Hour)
	case 5:
		dayTime = BeforeTime(-120 * time.Hour)
	case 6:
		dayTime = BeforeTime(-144 * time.Hour)
	case 7:
		dayTime = BeforeTime(-168 * time.Hour)
	default:
	}
	return dayTime
}

// 获取之前的日前
func BeforeTime(d time.Duration) string {
	return time.Now().Add(d).Format("2008-01-02")
}