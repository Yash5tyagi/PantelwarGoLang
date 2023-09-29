package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	mymap := make(map[string]int)

	ans := strings.Split(s, " ")
	// fmt.Println(ans)

	for i := 0; i < len(ans); i++ {
		mymap[ans(i)]++
	}

	return mymap
}

func main() {
	wc.Test(WordCount)
}
