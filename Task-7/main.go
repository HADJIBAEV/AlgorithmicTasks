package main

import "fmt"

func main() {
	var A, B, C, D int
	_, err := fmt.Scan(&A, &B, &C, &D)
	if err != nil {
		return
	}
	if D > B {
		s := A + C*(D-B)
		fmt.Println(s)
	} else {
		fmt.Println(A)
	}
}
