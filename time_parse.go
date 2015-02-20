package main

import (
	"fmt"
	"strings"
	"time"
)

func GetTimeFromString(time_string string) (int, int, int) {
	str_array := strings.Split(time_string, " ")
	var year, mon, day int
	fmt.Sscanf(str_array[0], "%d-%d-%d", &year, &mon, &day)
	return year, mon, day
}

func GetServerTime() string {
	t := time.Now()
	return fmt.Sprintf("%s",
		t.Local())
}

func CheckTimeIfExpired(purshed_date string, period int) bool {
	const time_short_format string = "01/02/2006"
	p_year, p_mon, p_day := GetTimeFromString(purshed_date)
	fmt.Println(p_year, p_mon, p_day)
	t_purchse, _ := time.Parse(time_short_format, fmt.Sprintf("%02d/%02d/%04d", p_mon, p_day, p_year))

	t_now := time.Now()
	fmt.Println("purchase", t_purchse)
	fmt.Println(t_now.After(t_purchse))
	if period == 100 {
		t_purchse = t_purchse.AddDate(0, 1, 0)
	} else if period == 600 {
		t_purchse = t_purchse.AddDate(0, 6, 0)

	} else if period == 1200 {
		t_purchse = t_purchse.AddDate(1, 0, 0)
	}

	fmt.Println("purchase changed:", t_purchse)
	return t_purchse.After(t_now)
}

func main() {
	p := fmt.Println

	fmt.Println("Hello, playground", GetServerTime())
	year, mon, day := GetTimeFromString(GetServerTime())
	fmt.Println(year, mon, day)

	CheckTimeIfExpired("2015-02-18 02:41:03 Etc/GM", 100)
	CheckTimeIfExpired("2015-02-18 02:41:03 Etc/GM", 600)
	CheckTimeIfExpired("2015-02-18 02:41:03 Etc/GM", 1200)
	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	//const shortForm = "2006-Jan-02"
	t := time.Now()
	t.Format("2006-01-02T15:04:05.999999-07:00")

	p("---This can parse for time --")
	//http://stackoverflow.com/questions/14106541/go-parsing-date-time-strings-which-are-not-standard-formats
	test, _ := time.Parse("01/02/2006", "02/19/2015")
	p(test)
	p(test.After(time.Now()))
	//t3 := t.Format("2006-01-02T15:04:05")

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)
	//fmt.Println(t2.Before(t1))
}
