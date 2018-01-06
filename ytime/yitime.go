package ytime

import (
	//"errors"
	bmath "billy/math"
	"billy/yi"
	//"math"
	"time"
)

//GetYiStringForDate 计算月份的天干地支序号
func GetYiStringForDate(t time.Time) (yheaven string, yearth string,
	mheaven string, mearth string,
	dheaven string, dearth string) {

	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	year0 := ZeroPoint.Year()

	//计算年份的天干地支

	var y int

	//立春在2月，立春后才是新一年，1984年对应“甲子”年，排1,数组序号对应0
	if month < 2 {
		y = year - year0 - 1
	} else {
		y = year - year0
	}

	yheaven, yearth = yi.EncodeHeavenAndEarth(y)

	//计算月份的天干地支,1984年1月对应“乙丑”排2，数组序号对应1
	//先计算当月第一节是几号
	fdayofsolar := GetDayForSolarTerm(year, month*2-2)

	//若未到当月第一气节，则是上个月
	m := (year-year0)*12 + month + 2 - 2
	m = bmath.If(day < fdayofsolar, m-1, m).(int)

	mheaven, mearth = yi.EncodeHeavenAndEarth(m)

	//计算当日的天干地支 "1984-01-10", "癸亥年 乙丑月 癸卯日",
	deltaD := int(t.Sub(ZeroPoint).Hours()) / 24
	deltaD = bmath.If(year < year0, deltaD-1, deltaD).(int)
	dayx := 36 + deltaD
	// fmt.Print(dayx)
	dheaven, dearth = yi.EncodeHeavenAndEarth(dayx)

	// y1 := year / 100
	// y2 := year % 100

	// tdh := 4*y1 + y1/4 + 5*y2 + y2/4 + 3*(month+1)/5 + day - 3
	// tde := tdh + 4*y1 + 10 + bmath.If((month%2) == 0, 6, 0).(int)

	// dh := tdh % 10
	// de := tde % 12

	// dheaven, dearth = yi.EncodeHeavenString(dh), yi.EncodeEarthString(de)

	return
}
