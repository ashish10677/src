package main

import "fmt"

func main() {
	no := 13
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d * %d = %d\n", no, i, (no * i))
	}
}
