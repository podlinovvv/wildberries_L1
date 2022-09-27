package main

import (
	"fmt"
	"strings"
)

func main() {
	st1 := "Qwerty"
	st2 := "Qqwerty"
	st3 := "qqwwerty"
	st4 := "    "
	st5 := ""
	fmt.Println(checkDoubles(&st1))
	fmt.Println(checkDoubles(&st2))
	fmt.Println(checkDoubles(&st3))
	fmt.Println(checkDoubles(&st4))
	fmt.Println(checkDoubles(&st5))

}

func checkDoubles(s *string) bool {
	lowcaseRunes:= []rune(strings.ToLower(*s))
	dict := make(map[rune]bool)
	for _,v := range lowcaseRunes {
		_, ok := dict[v]
		if ok {
			return false
		}
		dict[v] = true
	}
	return true
}
