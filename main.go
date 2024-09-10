package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Маппинг римских цифр на арабские
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Маппинг арабских цифр на римские
var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
}

// Основная функция калькулятора
func main() {
	// Чтение строки с консоли
	var input string
	fmt.Println("Введите выражение (например, 2 + 2 или VI / III):")
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
		num1, num2, err = convertRomanToArabic(operands[0], operands[1])
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
		fmt.Println(arabicToRoman[result])
	} else {
		fmt.Println(result)
	}
}

// Функция для разбора строки на операнды и оператор
func parseInput(input string) ([]string, string, error) {
	// Определение доступных операторов
	operators := []string{"+", "-", "*", "/"}

	// Поиск оператора в строке
	var operator string
	for _, op := range operators {
		if strings.Contains(input, op) {
			operator = op
			break
		}
	}

	// Если оператор не найден, это ошибка
	if operator == "" {
		return nil, "", errors.New("некорректная операция")
	}

	// Разделение строки по оператору
	operands := strings.Split(input, operator)
	if len(operands) != 2 {
		return nil, "", errors.New("некорректный формат операции")
	}

	// Убираем пробелы из операндов
	for i := range operands {
		operands[i] = strings.TrimSpace(operands[i])
	}

	return operands, operator, nil
}

// Функция для проверки, используются ли римские цифры
func isRomanNumerals(operands []string) bool {
	_, ok1 := romanToArabic[operands[0]]
	_, ok2 := romanToArabic[operands[1]]
	if ok1 && ok2 {
		return true
	} else if ok1 || ok2 {
		panic("используются одновременно римские и арабские цифры")
	}
	return false
}

// Функция для преобразования римских чисел в арабские
func convertRomanToArabic(op1, op2 string) (int, int, error) {
	num1, ok1 := romanToArabic[op1]
	num2, ok2 := romanToArabic[op2]
	if !ok1 || !ok2 {
		return 0, 0, errors.New("некорректные римские числа")
	}
	return num1, num2, nil
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
