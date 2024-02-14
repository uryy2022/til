package main

import "fmt"

var (
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	//	var t bool = true
	//	var f bool = false
	t, f bool = true, false
)

func main() {
	fmt.Println(i, f64, s, t, f)

	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)

}
