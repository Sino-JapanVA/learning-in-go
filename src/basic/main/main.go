package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	lrc, err := ioutil.ReadFile("../告白气球.txt")
	// txt := string(lrc)
	runeData := bytes.Runes(lrc)
	// runeTxt := string(runeData)
	checkErr(err)
	fmt.Println(len(lrc))
	fmt.Println(len(runeData))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
