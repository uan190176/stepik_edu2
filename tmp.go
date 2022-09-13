package main

import "fmt"

type iIter interface {
	getNext() bool
	getVal() int
}

type sIter struct {
	nums []int
	idx  int
}

func (it *sIter) getNext() bool {
	if it.idx+1 < len(it.nums) {
		return true
	}
	return false
}
func (it *sIter) getVal() int {
	defer func() { it.idx++ }()
	return it.nums[it.idx]
}

func newIter(src []int) *sIter {
	return &sIter{nums: src, idx: 0}
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	it := newIter(numbers)
	for it.getNext() {
		cur := it.getVal()
		fmt.Println(cur)
	}
}
