/*
Напишите программу, которая определяет название языка по его коду. Правила:

en → English
fr → French
ru или rus → Russian
иначе → Unknown
*/
package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// ...
	var lang string
	switch {
	case code == "en":
		lang = "English"
	case code == "fr":
		lang = "French"
	case code == "ru" || code == "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
