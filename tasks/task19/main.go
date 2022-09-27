package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	//превращаем текст в слайс рун, чтобы не было проблем с unicode-символами
	runeText := []rune(text)

	//переворачиваем слайс inplace
	for l, r := 0, len(runeText)-1; l < r; {
		runeText[l], runeText[r] = runeText[r], runeText[l]
		l++
		r--
	}

	fmt.Println(string(runeText))

}
