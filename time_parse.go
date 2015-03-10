package main

import (
	"fmt"
	"strings"
	"time"
)

func ParseCustomTimeString(time_string string) time.Time {
	//"2015-03-10 23:58:23 UTC"
	str_array := strings.Split(time_string, "UTC")
	var year, mon, day, hh, mm, ss int
	fmt.Sscanf(str_array[0], "%d-%d-%d %d:%d:%d ", &year, &mon, &day, &hh, &mm, &ss)
	time_string_to_parse := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00", year, mon, day, hh, mm, ss)
	url_create_time, _ := time.Parse(time.RFC3339, time_string_to_parse)
	return url_create_time
}

func CheckMapUrlExpired(time_string string) bool {
	//2015-03-08 21:40:54.370967455 +0800 CST
	str_array := strings.Split(time_string, ".")
	var year, mon, day, hh, mm, ss int
	fmt.Sscanf(str_array[0], "%d-%d-%d %d:%d:%d", &year, &mon, &day, &hh, &mm, &ss)
	time_string_to_parse := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00", year, mon, day, hh, mm, ss)
	//time_zone := time.Now().Location()
	//fmt.Println("time zone=", time_zone.String())
	url_create_time, _ := time.Parse(time.RFC3339, time_string_to_parse)
	url_expired_time := url_create_time.AddDate(0, 0, +1)
	fmt.Println("expired time:", url_expired_time, " current time:", time.Now())
	fmt.Println("Is is expired?", url_expired_time.Before(time.Now()))
	return url_expired_time.Before(time.Now())
}

func GetTimeFromString(time_string string) (int, int, int) {
	str_array := strings.Split(time_string, " ")
	var year, mon, day int
	fmt.Sscanf(str_array[0], "%d-%d-%d", &year, &mon, &day)
	return year, mon, day
}

func GetUTCTime() time.Time {
	t := time.Now()
	local_location, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
	}
	time_UTC := t.In(local_location)
	return time_UTC
}

func GetServerTime() string {
	return GetUTCTime().String()
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

//Import GMT transfer to local time.format
func TimeConvertGMT(gmt_time string) string {
	var year, mon, day, hour, minute, second int
	fmt.Sscanf(gmt_time, "%d-%d-%d %d:%d:%d Etc/GMT", &year, &mon, &day, &hour, &minute, &second)
	// fmt.Println(year, mon, day, hour, minute, second)
	time_format := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00", year, mon, day, hour, minute, second)
	// fmt.Println(time_format)

	t1, _ := time.Parse(time.RFC3339, time_format)
	time_string := t1.Local().String()
	// fmt.Println(time_string)

	//2012-11-02 06:08:41 +0800 CST
	var time_date, time_time, time_fix, time_zone string
	fmt.Sscanf(time_string, "%s %s %s %s", &time_date, &time_time, &time_fix, &time_zone)
	// fmt.Println(time_date, time_time, time_fix, time_zone)
	//2015-03-04 08:45:37.389639928 +0800 CST
	return fmt.Sprintf("%s %s.000000000 %s %s", time_date, time_time, time_fix, time_zone)
}

func main() {
	p := fmt.Println

	fmt.Println("Hello, playground", GetServerTime())
	year, mon, day := GetTimeFromString(GetServerTime())
	fmt.Println(year, mon, day)

	CheckTimeIfExpired("2015-02-18 02:41:03 Etc/GMT", 100)
	CheckTimeIfExpired("2015-02-18 02:41:03 Etc/GMT", 600)
	CheckTimeIfExpired("2015-02-18 02:41:03 Etc/GMT", 1200)
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

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	p(CheckMapUrlExpired(time.Now().String()))
	p(CheckMapUrlExpired("2015-03-08 11:48:54.832118324 +0800 CST"))

	//Change localtion to UTC and compare.
	time_local := time.Now()
	p("Local:", time_local)
	local_location, err := time.LoadLocation("UTC")
	if err != nil {
		p(err)
	}
	time_UTC := time_local.In(local_location)
	p("UTC:", time_UTC)
	p("XXX:", GetServerTime())
	// t1, _ := time.Parse(
	// 	time.RFC3339,
	// 	"2012-11-01T22:08:41+00:00")
	// p(t1)
	// p(t1.Local())
	// p(TimeConvertGMT("2015-03-04 09:59:59 Etc/GMT"))
	// p(time.Now().Local())
}
