/*
Программа принимает на вход строку source и число times. Требуется склеить source саму с собой times раз и вернуть результат:
source = x, times = 3 → xxx
source = omm, times = 2 → ommomm
*/
package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	// ...
	result := ""
	for i := 0; i < times; i++ {
		result += source
	}

	fmt.Println(result)
}
