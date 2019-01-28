package main

import "fmt"

func main() {
	if a := 9; a % 2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}