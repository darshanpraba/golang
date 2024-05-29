package main

import "fmt"

// this is a comment
var name string = "Dog rocky"

func main() {
	fmt.Println("Hello World")

	fmt.Println("1 + 1 =", 1+1)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x string
	x = `Hello World Darshan`
	fmt.Println(x)

	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	var x1 string = "hello"
	var y string = "world"
	fmt.Println(x1 == y)

	x2 := "Darshan Prabaharan"
	fmt.Println(x2)
	x2 = "Darshan Prabaharan Sivashanmugam"
	fmt.Println(x2)
	f(x2)

}

func f(x string) {
	fmt.Println(x)
	fmt.Println(name)
}
