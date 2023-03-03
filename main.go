package main

import (
	"fmt"
	"strconv"
	"strings"
)

const option = "+-*/"

func main() {

	var a, b, y string
	fmt.Scan(&a, &b, &y)
	fmt.Println(result(a, b, y))

}
func result(a, b, y string) string {
	parsSign(b)
	number1, num1 := strconv.Atoi(a)
	number2, num2 := strconv.Atoi(y)
	if num1 == nil && num2 == nil {
		return strconv.Itoa(parsNum(number1, number2, b))
	} else {
		return numberRome(a, b, y)
	}
}
func parsSign(b string) string {
	if strings.ContainsAny(b, option) {
		return b
	}
	panic("Неверный знак")
}
func parsNum(a, y int, b string) int {
	var res int
	if a < 11 && y < 11 {
		switch b {
		case "+":
			res = a + y
		case "-":
			res = a - y
		case "*":
			res = a * y
		case "/":
			res = a / y
		}
		return res
	} else {
		panic("Выберите только арабские либо римские числа <11")
	}
}
func numberRome(a, b, y string) string {
	numRome := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
		"VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	number1Rome, num1R := numRome[a]
	number2Rome, num2R := numRome[y]
	if num1R && num2R {
		num := parsNum(number1Rome, number2Rome, b)
		if num < 0 {
			panic("в римской системе отсуствуют отрицательные числа")
		}
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		var result strings.Builder
		for i := 0; i < len(symbols); i++ {
			for num >= value[i] {
				result.WriteString(symbols[i])
				num -= value[i]
			}

		}
		return result.String()
	} else {
		panic("Выберите только арабские либо римские числа <11")
	}

}
