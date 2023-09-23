package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var text string
	var waktu string
	if len(os.Args) > 1 {
		text = os.Args[1]
	} else {
		tNow := time.Now()
		tHour := tNow.Hour()
		if tHour >= 6 && tHour < 12 {
			waktu = "pagi"
		} else if tHour >= 12 && tHour < 18 {
			waktu = "siang"
		} else {
			waktu = "malam"
		}
		text = "selamat " + waktu
	}
	cwa(text)
}

func cwa(text string) {
	chars := strings.Split(text, "")
	charMap := make(map[string]int)
	for _, char := range chars {
		fmt.Println(char)
		charMap[char]++
	}
	fmt.Println()
	fmt.Println(charMap)
}
