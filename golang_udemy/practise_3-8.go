package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("init!")
}

func main() {
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
}
