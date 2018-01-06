package yi

import (
	"billy/math"
	// "fmt"
)

//EncodeHeavenString 将天干序号转化为天干文字
func EncodeHeavenString(index int) (heaven string) {
	i := math.If(index > -1, index%10, (index%10+10)%10).(int)
	// fmt.Println(i)
	heaven = HeavenlyStems[i]
	return
}

//EncodeEarthString 将地支序号转化为地支文字
func EncodeEarthString(index int) (earth string) {
	i := math.If(index > -1, index%12, (index%12+12)%12).(int)
	// fmt.Println(i)
	earth = EarthlyBranches[i]
	return
}

//EncodeHeavenAndEarth 转换将序号天干地支文字
func EncodeHeavenAndEarth(index int) (heaven, earth string) {
	heaven = EncodeHeavenString(index)
	earth = EncodeEarthString(index)
	return
}
