/*
Определите интерфейс универсального итератора (iterator), который можно использовать в функции выбора максимального элемента (max).
Реализуйте интерфейс для итератора по срезу целых чисел.
Подробности — по коду задания.
Если не понимаете, как подступиться к задаче, вот подсказка.

Sample Input:
1 4 5 2 3
Sample Output:
5
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	// чтобы понять сигнатуры методов - посмотрите,
	// как они используются в функции max() ниже
	next() bool
	val() element
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	// поля структуры
	nums []int
	idx  int
}

// методы intIterator, которые реализуют интерфейс iterator
func (it *intIterator) next() bool {
	if it.idx < len(it.nums) {
		return true
	}
	return false
}

func (it *intIterator) val() element {
	defer func() { it.idx++ }()
	return it.nums[it.idx]
}

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
	// ...
	return &intIterator{nums: src, idx: 0}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
	//nums := readInput2()
	nums := GetNumbers()
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput2() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if num != 0 {
			nums = append(nums, num)
		} else {
			break
		}
	}
	return nums
}

func GetNumbers() []int {
	return []int{1, 2, 3, 42}
}
