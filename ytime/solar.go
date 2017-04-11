package ytime

import (
	"time"
)

var solarTerm = []string{"小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "冬至"}
var sTermInfo = []int64{0, 21208, 42467, 63836, 85337, 107014, 128867, 150921, 173149, 195551, 218072, 240693, 263343, 285989, 308563, 331033, 353350, 375494, 397447, 419210, 440795, 462224, 483532, 504758}

//时间零点
var ZeroPoint = time.Date(1984, time.January, 6, 11, 40, 49, 00, time.UTC)

//GetDayForSolarTerm ===== 某年的第n个节气为几日(从0小寒起算)
func GetDayForSolarTerm(year, index int) (day int) {
	msoffset := int64(31556925974.7*float64(year-ZeroPoint.Year())) + sTermInfo[index]*60000
	offset := time.Duration(msoffset)
	newDate := ZeroPoint.Add(offset * time.Millisecond)
	day = newDate.Day()
	return
}
