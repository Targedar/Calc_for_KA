package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Маппинг римских цифр на арабские
var romanToArabic = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
}

// Маппинг для преобразования арабских чисел в римские
var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
	{1, "I"},
}

// Основная функция калькулятора
func main() {
	// Чтение строки с консоли
	var input string
	fmt.Println("Введите выражение (например, 2 + 2 или IX + VIII):")
	fmt.Scanln(&input)

	// Удаляем лишние пробелы
	input = strings.TrimSpace(input)

	// Разделение строки на операнды и оператор
	operands, operator, err := parseInput(input)
	if err != nil {
		panic(err.Error())
	}

	// Определение, римские или арабские числа используются
	isRoman := isRomanNumerals(operands)

	// Преобразование чисел в арабские для вычислений
	num1, num2 := 0, 0
	if isRoman {
		num1, err = romanToArabicConverter(operands[0])
		if err != nil {
			panic(err.Error())
		}
		num2, err = romanToArabicConverter(operands[1])
		if err != nil {
			panic(err.Error())
		}
	} else {
		num1, num2, err = convertArabic(operands[0], operands[1])
		if err != nil {
			panic(err.Error())
		}
	}

	// Выполнение арифметической операции
	result, err := calculate(num1, num2, operator)
	if err != nil {
		panic(err.Error())
	}

	// Вывод результата
	if isRoman {
		if result <= 0 {
			panic("В римской системе нет отрицательных чисел или нуля")
		}
		romanResult := arabicToRomanConverter(result)
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}

// Функция для разбора строки на операнды и оператор
func parseInput(input string) ([]string, string, error) {
	operators := []string{"+", "-", "*", "/"}

	var operator string
	for _, op := range operators {
		if strings.Contains(input, op) {
			operator = op
			break
		}
	}

	if operator == "" {
		return nil, "", errors.New("некорректная операция")
	}

	operands := strings.Split(input, operator)
	if len(operands) != 2 {
		return nil, "", errors.New("некорректный формат операции")
	}

	for i := range operands {
		operands[i] = strings.TrimSpace(operands[i])
	}

	return operands, operator, nil
}

// Функция для проверки, используются ли римские цифры
func isRomanNumerals(operands []string) bool {
	_, ok1 := romanToArabicConverter(operands[0])
	_, ok2 := romanToArabicConverter(operands[1])
	if ok1 == nil && ok2 == nil {
		return true
	} else if ok1 == nil || ok2 == nil {
		panic("используются одновременно римские и арабские цифры")
	}
	return false
}

// Функция для конвертации римских чисел в арабские
func romanToArabicConverter(roman string) (int, error) {
	roman = strings.ToUpper(roman)
	total := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value, exists := romanToArabic[string(roman[i])]
		if !exists {
			return 0, errors.New("некорректные римские числа")
		}

		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}
	return total, nil
}

// Функция для преобразования арабских чисел в римские
func arabicToRomanConverter(num int) string {
	result := ""
	for _, val := range arabicToRoman {
		for num >= val.Value {
			result += val.Symbol
			num -= val.Value
		}
	}
	return result
}

// Функция для преобразования арабских чисел
func convertArabic(op1, op2 string) (int, int, error) {
	num1, err1 := strconv.Atoi(op1)
	num2, err2 := strconv.Atoi(op2)
	if err1 != nil || err2 != nil || num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		return 0, 0, errors.New("некорректные арабские числа")
	}
	return num1, num2, nil
}

// Функция для выполнения арифметической операции
func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("некорректный оператор")
	}
}
