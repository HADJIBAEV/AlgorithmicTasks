package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	var s string
	var n int
	_, err := fmt.Fscan(in, &s)
	if err != nil {
		return
	}
	_, err = fmt.Fscan(in, &n)
	if err != nil {
		return
	}
	SliceS := strings.Split(s, "") // []
	for i := 0; i < n; i++ {
		var start, end int
		var stickers string
		_, err = fmt.Fscan(in, &start, &end, &stickers)
		if err != nil {
			return
		}
		Stick := strings.Split(stickers, "") // []
		for j := start; j <= end; j++ {
			SliceS[j-1] = Stick[j-start]
		}
	}
	_, err = fmt.Fprintln(out, strings.Join(SliceS, ""))
	if err != nil {
		return
	}
	err = out.Flush()
	if err != nil {
		return
	}
}
