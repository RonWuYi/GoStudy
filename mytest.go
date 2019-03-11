//package awesomeProject
package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y

}

func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python int = 3, 4
var java bool = true

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("my favorite is %g\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 45))
	a, b := swap("hello", "woo")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	k := 3
	fmt.Println(i, c, python, java, k)
}
