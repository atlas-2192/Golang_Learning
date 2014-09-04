package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
)

func main() {
	fname := "sample1.jpg"

	f, err := os.Open(fname)
	fmt.Println("start to open image")
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	fmt.Println("start to decode image")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(x.String())
	fmt.Println("start to get exif image")
	camModel, _ := x.Get(exif.Model)
	if camModel != nil {
		fmt.Println(camModel.StringVal())
	}

	fmt.Println("start to get exif - date from image")
	date, _ := x.Get(exif.DateTimeOriginal)
	if date != nil {
		fmt.Println(date.StringVal())
	}

	color_space, _ := x.Get(exif.ColorSpace)
	if color_space != nil {
		fmt.Println("Color Space Get..")
		fmt.Printf("Type is %T\n", color_space)
		fmt.Printf("oth data %d\n", color_space.TypeCategory())
	}
	fmt.Println("Done...")
}
