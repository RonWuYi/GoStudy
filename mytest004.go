package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "不变"
	v := 42
	fmt.Println("hello", World)
	fmt.Printf("v is of type %T\n", v)
	const Truth = true
	fmt.Println("go rules", Truth)
	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
