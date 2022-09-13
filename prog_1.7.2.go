/*
Я написал программу, которая считывает исходную строку вида:
balance/overdraft t1 t2 t3 ... tn
И превращает ее в счет типа account и список транзакций типа []int.
К сожалению, я поленился явно обработать ошибки, положившись на отлов паники через defer и recover(). Исправьте программу, заменив их на явный возврат и обработку ошибок.
В получившейся программе не должно быть конструкций defer и recover(). Если в исходной строке есть ошибка — программа должна вывести только ошибку. Иначе — вывести счет и транзакции.

Например:
на входе 80/10 10 -20 30
на выходе -> {80 10} [10 -20 30]
Или:
на входе 80/10 10 z 30
на выходе -> strconv.Atoi: parsing "z": invalid syntax
Если в исходной строке несколько ошибок, программа должна вывести только первую.
Не меняйте текст ошибок, используйте те же строки, что принимали вызовы panic() в исходной программе.
Гарантируется, что общий формат исходной строки соблюдается (то есть значения разделены пробелами, овердрафт отделен от баланса дробью, и тому подобное), но могут быть ошибки в отдельных значениях (как в примере выше).
Программа не должна завершаться паникой или os.Exit(1) ни при каких обстоятельствах.

Sample Input:
80/10 10 -20 30
Sample Output:
-> {80 10} [10 -20 30]
*/
package main

// не меняйте импорты, они нужны для проверки
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errOverdraft error = errors.New("expect overdraft >= 0")
var errBalance error = errors.New("balance cannot exceed overdraft")

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	var err error

	acc, trans, err = parseInput()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(acc, trans)
	}
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInputB()
	acc, err := parseAccount(accSrc)
	if err != nil {
		return account{}, nil, err
	}
	trans, err := parseTransactions(transSrc)
	if err != nil {
		return account{}, nil, err
	}

	return acc, trans, err
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInputB() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		if scanner.Text() == "0" {
			break
		}
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, _ := strconv.Atoi(parts[0])
	overdraft, _ := strconv.Atoi(parts[1])
	if overdraft < 0 {
		//panic("expect overdraft >= 0")
		return account{}, errOverdraft
	}
	if balance < -overdraft {
		//panic("balance cannot exceed overdraft")
		return account{}, errBalance
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, e := strconv.Atoi(s)
		if e == nil {
			trans[idx] = t
		} else {
			return trans, e
		}
	}
	return trans, nil
}
