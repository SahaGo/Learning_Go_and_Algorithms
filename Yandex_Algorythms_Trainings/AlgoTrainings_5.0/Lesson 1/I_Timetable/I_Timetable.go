package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type date struct {
	day        int
	month      string
	monthNum   int
	weekdayNum int
}

type weekday struct {
	name          string
	count         int
	holidaysCount int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var holNum, year, day int
	var holMonth, startWeekday string
	fmt.Fscan(in, &holNum, &year)

	arr := make([]date, holNum)

	for i := 0; i < holNum; i++ {
		fmt.Fscan(in, &day, &holMonth)
		arr[i] = date{day: day, month: holMonth}
	}
	fmt.Fscan(in, &startWeekday)

	days := []weekday{
		{name: "Sunday", count: 52},
		{name: "Monday", count: 52},
		{name: "Tuesday", count: 52},
		{name: "Wednesday", count: 52},
		{name: "Thursday", count: 52},
		{name: "Friday", count: 52},
		{name: "Saturday", count: 52}}

	isLeapYear := leap(year)

	for i := 0; i < 7; i++ {
		if isLeapYear {
			if days[i].name == startWeekday {
				days[i].count = days[i].count + 1
				if i == 6 {
					days[0].count = days[0].count + 1
				} else {
					days[i+1].count = days[i+1].count + 1
				}
				break
			}
		} else {
			if days[i].name == startWeekday {
				days[i].count = days[i].count + 1
				break
			}
		}
	}

	for i := 0; i < holNum; i++ {
		arr[i].weekdayNum = procDate(arr[i], year)
		days[arr[i].weekdayNum].holidaysCount++
	}

	maxDays := math.MinInt
	minDays := math.MaxInt
	maxInd := 0
	minInd := 0

	for i := 0; i < 7; i++ {
		sumHolidays := 0
		for j := 0; j < 7; j++ {
			if j == i {
				continue
			}
			sumHolidays += days[j].holidaysCount
		}

		if days[i].count+sumHolidays <= minDays {
			minDays = days[i].count + sumHolidays
			minInd = i
		}
		if days[i].count+sumHolidays >= maxDays {
			maxDays = days[i].count + sumHolidays
			maxInd = i
		}
	}

	//fmt.Fprintln(out, arr, "\n", days, "\n", days[maxInd].name, days[minInd].name)
	fmt.Fprintln(out, days[maxInd].name, days[minInd].name)
}

func leap(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func procDate(day date, year int) int {
	switch day.month {
	case "January":
		day.monthNum = 1
	case "February":
		day.monthNum = 2
	case "March":
		day.monthNum = 3
	case "April":
		day.monthNum = 4
	case "May":
		day.monthNum = 5
	case "June":
		day.monthNum = 6
	case "July":
		day.monthNum = 7
	case "August":
		day.monthNum = 8
	case "September":
		day.monthNum = 9
	case "October":
		day.monthNum = 10
	case "November":
		day.monthNum = 11
	case "December":
		day.monthNum = 12
	}

	if day.monthNum < 3 {
		year--
		day.monthNum += 10
	} else {
		day.monthNum -= 2
	}

	return ((day.day + 31*day.monthNum/12 + year + year/4 - year/100 + year/400) + 7) % 7
}
