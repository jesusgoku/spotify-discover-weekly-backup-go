package main

import (
	"testing"
	"time"
)

type testCase struct {
	input  time.Time
	expect string
}

func TestGetFirstDayOfWeek(t *testing.T) {
	testCases := []testCase{
		testCase{
			input:  time.Date(2019, time.December, 8, 23, 0, 0, 0, time.UTC),
			expect: "20191202",
		},
		testCase{
			input:  time.Date(2009, time.October, 7, 23, 0, 0, 0, time.UTC),
			expect: "20091005",
		},
		testCase{
			input:  time.Date(2009, time.September, 4, 23, 0, 0, 0, time.UTC),
			expect: "20090831",
		},
		testCase{
			input:  time.Date(1991, time.December, 15, 23, 0, 0, 0, time.UTC),
			expect: "19911209",
		},
		testCase{
			input:  time.Date(1985, time.December, 18, 23, 0, 0, 0, time.UTC),
			expect: "19851216",
		},
	}

	for _, currentTestCase := range testCases {
		got := GetFirstDayOfWeek(currentTestCase.input).Format("20060102")

		if got != currentTestCase.expect {
			t.Errorf("GetFirstDayOfWeek <- %+v = %+v; want %+v", currentTestCase.input, got, currentTestCase.expect)
		}
	}
}
