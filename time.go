package main

import (
	"fmt"
	"strings"
	"time"
)

func GetServerTime() string {
	//2015-02-04 09:10:51.728001775 +0800 CST
	t := time.Now()
	return fmt.Sprintf("%s",
		t.Local())
}

func GetTimeFromString(time_string string) (int, int, int) {
	str_array := strings.Split(time_string, " ")
	var year, mon, day int
	fmt.Sscanf(str_array[0], "%d-%d-%d", &year, &mon, &day)
	return year, mon, day
}

func main() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, time.Now().Format(time.RFC850), loc)
	fmt.Println(t)

	fmt.Println(time.Now().Format(time.RFC850))

	fmt.Println(GetServerTime())

	yy, mm, dd := GetTimeFromString(GetServerTime())
	fmt.Printf("%d-%d-%d \n", yy, mm, dd)
}
