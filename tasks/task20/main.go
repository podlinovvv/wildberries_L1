package main

import (
	"fmt"
	"strings"
)

func main() {
	st := "snow dog sun"
	st2 := "     snow   dog sun"
	st3 := "snow dog   sun     "
	reverse(st)
	reverse(st2)
	reverse(st3)
}

//это версия программы без использования split и сохраняет количество пробелов между словами

func reverse(input string) {
	i := []rune(input)
	var sb strings.Builder
	//с помощью двух указателей l и r находятся границы каждого слова с конца до начала строки
	for l, r := len(i)-1, len(i)-1; l >= 0; l-- {
		switch {
		//для самого первого слова
		case l == 0:
			sb.WriteString(string(i[l:r]))
		//для самого последнего слова
		case i[l] == ' ' && r == len(i)-1:
			sb.WriteString(string(i[l+1 : r+1]))
			sb.WriteString(" ")
			r = l
		//копирование идущих подряд пробелов
		case i[l] == ' ' && (l+1 == r):
			sb.WriteString(" ")
			r--
		//одиночный пробел
		case i[l] == ' ':
			sb.WriteString(string(i[l+1 : r+1]))
			r = l
		}
	}
	fmt.Println(sb.String())
}
