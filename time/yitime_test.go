package time

import (
	"fmt"
	"testing"
	"time"
)

type HEData struct {
	Date           string
	HeavenAndEarth string
	HAndE          string
}

var src = []HEData{
	HEData{"1984-01-10", "癸亥年 乙丑月 癸卯日", ""},
	HEData{"1985-01-10", "甲子年 丁丑月 己酉日", ""},
	HEData{"2010-10-07", "庚寅年 乙酉月 庚寅日", ""},
	HEData{"1904-09-07", "甲辰年 壬申月 甲辰日", ""},
	HEData{"1998-05-08", "戊寅年 丁巳月 乙卯日", ""},
	HEData{"2003-10-08", "癸未年 辛酉月 甲寅日", ""},
	HEData{"2009-07-08", "己丑年 辛未月 甲寅日", ""},
	HEData{"2015-04-23", "乙未年 庚辰月 己巳日", ""},
}

func TestYiDate(t *testing.T) {
	for _, d := range src {
		date, err := time.Parse("2006-01-02", d.Date)
		if err != nil {
			t.Error(err)
		}
		yh, ye, mh, me, dh, de := GetYiStringForDate(date)
		d.HAndE = yh + ye + "年 " + mh + me + "月 " + dh + de + "日"
		if d.HAndE != d.HeavenAndEarth {
			t.Error("不合:", d.Date, "-", d.HeavenAndEarth, "-", d.HAndE)
			//t.Fail()
		} else {
			fmt.Println("相合:", d.Date, "-", d.HeavenAndEarth, "-", d.HAndE)

		}
	}
}
