package main

import (
	"fmt"
	"strings"
)

func main() {
	// lrc, err := ioutil.ReadFile("../告白气球.txt")
	// // txt := string(lrc)
	// runeData := bytes.Runes(lrc)
	// // runeTxt := string(runeData)
	// checkErr(err)
	// fmt.Println(len(lrc))
	// fmt.Println(len(runeData))
	lrc := "ss超凡蜘蛛侠"
	lrcRune := []rune(lrc)
	txt := "蜘"
	txtRune := []rune(txt)
	n := strings.IndexRune(lrc, txtRune[0])
	fmt.Println(n)
	for index, val := range lrcRune {
		fmt.Printf("%v  %d\n", string(val), index)
	}
	fmt.Println(string(lrcRune[4]))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
