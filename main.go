package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result, err := calculate(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func calculate(input string) (string, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return "", errors.New("формат математической операции не удовлетворяет заданию")
	}

	num1Str, operator, num2Str := parts[0], parts[1], parts[2]
	isRoman := false
	var num1, num2 int
	var err error

	if isRomanNumeral(num1Str) && isRomanNumeral(num2Str) {
		isRoman = true
		num1 = romanNumerals[num1Str]
		num2 = romanNumerals[num2Str]
	} else if isArabicNumeral(num1Str) && isArabicNumeral(num2Str) {
		num1, err = strconv.Atoi(num1Str)
		if err != nil {
			return "", errors.New("ошибка преобразования первого числа")
		}
		num2, err = strconv.Atoi(num2Str)
		if err != nil {
			return "", errors.New("ошибка преобразования второго числа")
		}
	} else {
		return "", errors.New("используются одновременно разные системы счисления")
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		return "", errors.New("числа должны быть в диапазоне от 1 до 10 включительно")
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("деление на ноль невозможно")
		}
		result = num1 / num2
	default:
		return "", errors.New("неизвестный оператор")
	}

	if isRoman {
		if result < 1 {
			return "", errors.New("в римской системе нет отрицательных чисел")
		}
		return arabicToRoman[result], nil
	}

	return strconv.Itoa(result), nil
}

func isRomanNumeral(s string) bool {
	_, exists := romanNumerals[s]
	return exists
}

func isArabicNumeral(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
