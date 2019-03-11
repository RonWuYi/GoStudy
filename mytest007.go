package main

import (
	//"expvar"
	"fmt"
	"math"
)

func powx(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(
		powx(3, 2, 10),
		powx(3, 3, 20),
	)
	//fmt.Println("lim is", lim)
}
