package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	step := 1
	N := 0
	for n > step {
		step *= 2
		N++
	}
	fmt.Println(N)
}
