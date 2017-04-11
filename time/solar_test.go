package time

import (
	"fmt"
	"testing"
)

type testSolar struct {
	Year  int
	Index int
	Day   int
}

var solarlist = []testSolar{
	testSolar{1901, 0, 6},
	testSolar{1900, 1, 20},
	testSolar{2000, 1, 21},
}

func TestGetSolar(t *testing.T) {

	for _, s := range solarlist {
		day := GetDayForSolarTerm(s.Year, s.Index)

		fmt.Println(s.Year, ",", s.Index, ",", s.Day, "--", day)
		if day != s.Day {
			t.Fail()
		}
	}
}
