/*
Напишите программу, которая проверяет корректность пароля. Корректным считается пароль, который удовлетворяет любому из условий:
содержит буквы и цифры;
длина не менее 10 символов.
Следуйте указаниям по тексту программы. Не меняйте сигнатуры функций, определение типа password и переменную validator в main().
Гарантируется, что на вход программы подается строка без пробелов.

Sample Input:
hellowor1d
Sample Output:
true
*/
package main

import (
	"fmt"
	"unicode"
)

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func digits(s string) bool {
	// ...
	r := []rune(s)
	for _, c := range r {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func letters(s string) bool {
	// ...
	r := []rune(s)
	for _, c := range r {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false

}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func minlen(length int) validator {
	// ...
	return func(s string) bool {
		if len([]rune(s)) > 9 {
			return true
		} else {
			return false
		}
	}
}

// and возвращает валидатор, который проверяет, что все
// переданные ему валидаторы вернули true
func and(funcs ...validator) validator {
	// ...
	return func(s string) bool {
		for _, f := range funcs {
			if f(s) == false {
				return false
			}
		}
		return true
	}
}

// or возвращает валидатор, который проверяет, что хотя бы один
// переданный ему валидатор вернул true
func or(funcs ...validator) validator {
	// ...
	return func(s string) bool {
		for _, f := range funcs {
			if f(s) == true {
				return true
			}
		}
		return false
	}

}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	// ...
	return p.validator(p.value)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	var s string
	fmt.Scan(&s)
	// валидатор, который проверяет, что пароль содержит буквы и цифры,
	// либо его длина не менее 10 символов
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
