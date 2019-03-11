package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum is ", sum)
	//for ; ;  {
	//	fmt.Println(sum)
	//}
	fmt.Println(sqrt(2), sqrt(-4))
}
