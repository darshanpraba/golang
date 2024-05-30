package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/StephaneBunel/bresenham"
)

func main() {
	const x string = "Hello World"
	fmt.Println(x)
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1, slice2)

	var imgRect = image.Rect(0, 0, 500, 500)
	var img = image.NewRGBA(imgRect)
	var colBLUE = color.RGBA{0, 0, 255, 255}

	// draw line
	bresenham.DrawLine(img, 14, 71, 441, 317, colBLUE)

	// save image in example1.png
	toimg, _ := os.Create("example1.png")
	defer toimg.Close()
	png.Encode(toimg, img)
}
