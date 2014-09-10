package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//ext_str := fmt.Sprintf("&#%d;&#%d;", '世', '界')
	ext_str := fmt.Sprintf("%q", "世界")
	fmt.Println(ext_str)

	// output, err := exec.Command(`d:\exiftool.exe`,
	// 	`-title="test 2 世界"`,
	// 	//If it not trasnfer from 世界 -> &#19990;&#30028;  it could not display correctly.
	// 	"-E", "test.jpg").CombinedOutput()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(string(output))

	str_title := fmt.Sprintf("-title=\"%s\"", ext_str)
	stdout, _ := exec.Command(`d:\exiftool.exe`, str_title, "-E", "test.jpg").CombinedOutput()
	fmt.Println(string(stdout))
}
