/*
Напишите функцию filter(), которая фильтрует срез целых чисел с помощью функции-предиката и возвращает отфильтрованный срез.
Функция-предикат вызывается для каждого элемента исходного среза. Если она возвращает true, элемент попадает в отфильтрованный срез.
Если возвращает false — не попадает.
Считайте исходный срез из стандартного ввода с помощью готовой функции readInput(). Затем выполните на нем filter().
В качестве предиката используйте функцию, которая возвращает true только для четных чисел. Напечатайте отфильтрованный срез.
Гарантируется, что на вход подаются только целые числа.

Sample Input:
1 2 3 4 5 6
Sample Output:
[2 4 6]
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
	var res []int
	for _, i := range iterable {
		if predicate(i) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	src := readInput()
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	res := filter(func(a int) bool { return bool(a%2 == 0) }, src)
	fmt.Println(res)
	//fmt.Printf("%T", src)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
		if num == 0 {
			break
		}
	}
	return nums
}
