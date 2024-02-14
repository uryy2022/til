package main

import (
	"fmt"
	"golang-test/hello"
)

/*
import (
	"fmt"
	"golang-test/hello"
)

func main() {
	greeting, err := hello.Hello("World !", "Spanish")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}

	greeting, err = hello.Hello("", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}
	greeting, err = hello.Hello("", "Russian")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}
}
*/

func main() {
	fmt.Println(hello.Hello("", ""))
}
