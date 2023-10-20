package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var NumberOfCard int
	_, err := fmt.Scan(&NumberOfCard)
	if err != nil {
		return
	}

	Subsequence := make([]int, NumberOfCard)
	WinningSequence := make([]int, NumberOfCard)

	for v := range Subsequence { // a1,a2, ... an
		_, err2 := fmt.Scan(&Subsequence[v])
		if err2 != nil {
			return
		}
	}

	for v := range WinningSequence { //b1,b2, ... bn
		_, err2 := fmt.Scan(&WinningSequence[v])
		if err2 != nil {
			return
		}
	}

	if CheckSlicesEquality(Subsequence, WinningSequence) {
		fmt.Println("YES")
		return
	}

	LeftBound, RightBound := FindLeftRightBound(NumberOfCard, Subsequence, WinningSequence) // Find Left Bound & Right Bound

	sort.Ints(Subsequence[LeftBound : RightBound+1]) //sort Left Bound & Right Bound

	if CheckSlicesEquality(Subsequence, WinningSequence) { // Subsequence == WinningSequence
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// FindLeftRightBound - Find Left bound & Right bound
func FindLeftRightBound(n int, Subsequence, WinningSequence []int) (LeftBound int, RightBound int) {
	for i := 0; i < n; i++ {
		if Subsequence[i] != WinningSequence[i] {
			if LeftBound == 0 {
				LeftBound = i
			} else {
				RightBound = i
			}
		}
	}
	return LeftBound, RightBound
}

// CheckSlicesEquality - Check equality of slices
func CheckSlicesEquality(Subsequence, WinningSequence []int) bool {
	return reflect.DeepEqual(Subsequence, WinningSequence) // Subsequence == WinningSequence
}
