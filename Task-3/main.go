package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Database struct {
	m map[string]string
}

func NewDatabase() *Database {
	return &Database{
		m: make(map[string]string),
	}
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ConvertRune := []rune(str)
	if unicode.IsUpper(ConvertRune[0]) && string(ConvertRune[len(ConvertRune)-3]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
	//unicode.IsUpper(ConvertRune[0]) &&
	//var str, oftenStr string
	//fmt.Scan(&str)
	//var count int
	//k := make(map[string]int)
	//for v := range str {
	//	if v != count {
	//		k = append(k, v)
	//		k[v] = 0
	//	}
	//}
}
