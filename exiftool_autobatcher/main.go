package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type visitor int

func walkpath(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s with %d bytes\n", path, f.Size())
	if CaseInsensitiveContains(filepath.Ext(path), "JPG") {
		fmt.Println("It is image files........")
		//Read 1.jpg.txt
		data_file_name := fmt.Sprintf("%s.txt", path)
		fmt.Printf("Check %s........\n", data_file_name)

		dat, err := ioutil.ReadFile(data_file_name)
		if err == nil {
			fmt.Println(string(dat))
			fmt.Printf("&#%d;\n", string(dat))
			ext_str := fmt.Sprintf("&#%d;", '許')
			fmt.Println(ext_str)
			str_abs, _ := filepath.Abs(path)
			//cmd_str := fmt.Sprintf("cmd /c exiftool.exe -title=\"1234\"  %s", str_abs)
			//fmt.Println(string(cmd_str))

			str_title := fmt.Sprintf("-title=\"%s\"a", ext_str)
			fmt.Println(ext_str)
			stdout, err := exec.Command("cmd", "/c", "exiftool.exe", str_title, "-E", str_abs).Output()
			fmt.Println(string(stdout))
			if err != nil {
				fmt.Printf("Run command error: %s\n", err.Error())
			}
		} else {
			fmt.Printf("%s\n", err.Error())
		}
	}
	return nil
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func RunExifTool(source_file, title string) {
	request_url := "http://google.com"
	var err error
	switch runtime.GOOS {
	case "linux":
		fmt.Println("Mac/Linux")
		err = exec.Command("xdg-open", request_url).Start()
	case "windows", "darwin":
		fmt.Println("Windows")
		fmt.Println(request_url)
		err = exec.Command("cmd", "/c", "start", request_url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println(err)
	}
}

//type WalkFunc func(path string, info os.FileInfo, err error) error
func main() {
	// Check exiftool if exit in this tool.
	exec_file_name := "exiftool.exe"
	if _, err := os.Stat(exec_file_name); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", exec_file_name)
		return
	}

	flag.Parse()
	root := flag.Arg(0)
	if len(root) != 0 {
		filepath.Walk(root, walkpath)
	}
}
