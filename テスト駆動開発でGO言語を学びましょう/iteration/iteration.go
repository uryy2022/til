package main

import "fmt"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
		//		fmt.Println(repeated)
	}
	return repeated
}

func main() {
	Repeats := Repeat("a")
	fmt.Printf("result=%q", Repeats)
}
