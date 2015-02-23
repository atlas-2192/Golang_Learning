// Sometimes we'll want to sort a collection by something
// other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts
// in Go.

package main

import (
	"fmt"
	"sort"
	"strings"
)

type DisplayCountTable struct {
	IndexName string
	Count     int
}

// In order to sort by a custom function in Go, we need a

func GetTimeFromString(time_string string) (int, int, int) {
	str_array := strings.Split(time_string, " ")
	var year, mon, day int
	fmt.Sscanf(str_array[0], "%d-%d-%d", &year, &mon, &day)
	return year, mon, day
}

// In order to sort by a custom function in Go, we need a
// corresponding type. Here we've created a `ByLength`
// type that is just an alias for the builtin `[]string`
// type.

type SortByTime []DisplayCountTable

// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. `Len` and `Swap`
// will usually be similar across types and `Less` will
// hold the actual custom sorting logic. In our case we
// want to sort in order of increasing string length, so
// we use `len(s[i])` and `len(s[j])` here.

func (s SortByTime) Len() int {
	return len(s)
}
func (s SortByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortByTime) Less(i, j int) bool {
	year, mon, day := GetTimeFromString(s[i].IndexName)
	s_i := year*10000 + mon*100 + day

	// fmt.Println("s_i:", s_i)

	year, mon, day = GetTimeFromString(s[j].IndexName)
	s_j := year*10000 + mon*100 + day
	// fmt.Println("s_j:", s_j)

	return s_i < s_j
}

// With all of this in place, we can now implement our
// custom sort by casting the original `fruits` slice to
// `ByLength`, and then use `sort.Sort` on that typed
// slice.
func main() {

	time_slice := make([]DisplayCountTable, 0)
	time_slice = append(time_slice, DisplayCountTable{IndexName: "2015-11-23", Count: 3})
	time_slice = append(time_slice, DisplayCountTable{IndexName: "2014-12-03", Count: 2})
	time_slice = append(time_slice, DisplayCountTable{IndexName: "2014-11-25", Count: 1})
	time_slice = append(time_slice, DisplayCountTable{IndexName: "2016-10-25", Count: 8})
	fmt.Println("Before sort:", time_slice)
	sort.Sort(SortByTime(time_slice))
	fmt.Println("After sort:", time_slice)

	fmt.Println("Less sort:", SortByTime(time_slice).Less(1, 2))
}
