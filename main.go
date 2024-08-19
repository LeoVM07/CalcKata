package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	symbol := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := ""
	for i := 0; num > 0; i++ {
		for num >= value[i] {
			num -= value[i]
			str += symbol[i]
		}
	}
	return str
}
func RomanCheck(RN []string, rn1 string) int {
	for _, rc := range RN {
		if rc == rn1 {
			return 1
		}
	}
	return 0
}

func IfArabic(n string) int {
	IR := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, ir := range IR {
		if ir == n {
			return 1
		}
	}
	return 0
}
func ArabicCalc(na1, na2 int, oprt string) {
	if na1 > 10 {
		println("На входе больше 10!")
	} else if na2 > 10 {
		println("На входе больше 10!")
	} else {
		switch oprt {
		case "+":
			fmt.Println(na1 + na2)
		case "-":
			fmt.Println(na1 - na2)
		case "*":
			fmt.Println(na1 * na2)
		case "/":
			fmt.Println(na1 / na2)
		default:
			fmt.Println("Do it again!")
		}
	}
}

func RomanCalc(rn1, rn2, oprt string) {
	RomanNumbers := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	RomanConversion := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	RC := RomanCheck(RomanNumbers, rn1) + RomanCheck(RomanNumbers, rn2)
	if RC == 2 {
		switch oprt {
		case "+":
			sum := RomanConversion[rn1] + RomanConversion[rn2]
			println(intToRoman(sum))
		case "-":
			sum := RomanConversion[rn1] - RomanConversion[rn2]
			if sum <= 0 {
				fmt.Println("Результат меньше или равен нулю!")
			} else {
				println(intToRoman(sum))
			}
		case "*":
			sum := RomanConversion[rn1] * RomanConversion[rn2]
			println(intToRoman(sum))
		case "/":
			sum := RomanConversion[rn1] / RomanConversion[rn2]
			println(intToRoman(sum))
		default:
			fmt.Println("Некорректное выражение!")
		}
	} else {
		fmt.Println("Некорректное выражение!")
	}
}

func main() {
	fmt.Println("Ведите выражение с арабскими или римскими цифрами")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		output := strings.Fields(input)
		n1, oprt, n2 := output[0], output[1], output[2]
		if len(output) == 3 {
			AC := IfArabic(n1) + IfArabic(n2)
			if AC == 2 {
				n1, _ := strconv.Atoi(n1)
				n2, _ := strconv.Atoi(n2)
				ArabicCalc(n1, n2, oprt)
			} else {
				RomanCalc(n1, n2, oprt)
			}
		} else {
			fmt.Println("Неверное количество переменных!")
		}
	}

}
