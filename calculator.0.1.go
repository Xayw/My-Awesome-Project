package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ERR1 = "ОШИБКА! Строка не является математической операцией."
	ERR2 = "ОШИБКА! Формат математических операций не соответствует требованию."
	ERR3 = "ОШИБКА! Одновременно используются разные системы счисления."
	ERR4 = "ОШИБКА! В римской системе нет отрицательных чисел."
	ERR5 = "ОШИБКА! В римской системе нет числа 0."
	ERR6 = "ОШИБКА! Калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно"
)

var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

func main() {
	fmt.Println("Введите выражение в формате 'число_оператор_число': ")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		var str = strings.ReplaceAll(console, " ", "")
		str = strings.TrimSpace(str)
		str = strings.ToUpper(str)
		base(str)
	}
}

var data []string

func base(s string) {
	var operator string
	var stringsFound int
	arabicNum := make([]int, 0)
	romansNum := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic(ERR2)
	case len(operator) < 1:
		panic(ERR1)
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romansNum = append(romansNum, elem)
		} else {
			arabicNum = append(arabicNum, num)
		}
	}
	switch stringsFound {
	case 0:
		errCheck := arabicNum[0] > 0 && arabicNum[0] < 11 && arabicNum[1] > 0 && arabicNum[1] < 11
		if val, ok := operators[operator]; ok && errCheck == true {
			a, b = &arabicNum[0], &arabicNum[1]
			fmt.Println(val())
		} else {
			panic(ERR6)
		}
	case 1:
		panic(ERR3)
	case 2:
		for _, elem := range romansNum {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(ERR6)
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(ERR5)
	} else if romanResult < 0 {
		panic(ERR4)
	}
	var convIntRoman = [...]int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for romanResult > 0 {
		for _, elem := range convIntRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Printf(romanNum)
}
