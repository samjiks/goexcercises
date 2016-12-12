package main

import "fmt"

func multiply(i int) int {
	return i * i
}

func params(a int, b string) (x int, y string) {
	x = a
	y = b
	return x, y
}

func naked(s string) (ss string) {
	ss = s
	return
}

func main() {
	fmt.Println(multiply(10))
	fmt.Println(params(1, "one"))
	fmt.Println(naked("returning naked string"))
}
