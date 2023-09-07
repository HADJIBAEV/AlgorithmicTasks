package main

import "fmt"

func main() {
	var source string
	const simple = "sheriff"
	_, err := fmt.Scan(&source)
	if err != nil {
		return
	}
	LettersSource := make(map[rune]int)
	LettersSample := make(map[rune]int)
	for _, v := range source {
		LettersSource[v]++
	}
	for _, v := range simple {
		LettersSample[v]++
	}
	minResult := 200000 // 2 * 10^5
	for i, v := range LettersSample {
		var division = LettersSource[i] / v
		if division < minResult {
			minResult = division
		}
	}
	fmt.Println(minResult)
}
