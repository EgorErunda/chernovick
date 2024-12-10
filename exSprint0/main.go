package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// ошибка, возвращаемая в случае некорректного выражения.
var ErrInvalidExpression = errors.New("invalid expression")

// вычисляет значение строкового выражения
func Calc(expression string) (float64, error) {
	// Удаляем все пробелы из выражения
	expression = removeSpaces(expression)

	// Проверяем баланс скобок
	if !isBalanced(expression) {
		return 0, ErrInvalidExpression
	}

	// Преобразуем выражение в обратную польскую нотацию
	rpn, err := shuntingYard(expression)
	if err != nil {
		return 0, err
	}

	// Вычисляем значение выражения
	result, err := evaluateRPN(rpn)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// удаляет все пробелы из строки
func removeSpaces(s string) string {
	var result string
	for _, char := range s {
		if !unicode.IsSpace(char) {
			result += string(char)
		}
	}
	return result
}

// проверяет кол-во скобок в строке
func isBalanced(s string) bool {
	count := 0
	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// преобразование выражения в обратную польскую нотацию
func shuntingYard(expression string) ([]interface{}, error) {
	var output []interface{}
	var operators []rune

	for i := 0; i < len(expression); {
		char := rune(expression[i])

		if unicode.IsDigit(char) || char == '.' {
			// Читаем число
			var number string
			for ; i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.'); i++ {
				number += string(expression[i])
			}
			num, err := strconv.ParseFloat(number, 64)
			if err != nil {
				return nil, ErrInvalidExpression
			}
			output = append(output, num)
		} else if isOperator(char) {
			// Обрабатываем оператор
			for len(operators) > 0 && isOperator(operators[len(operators)-1]) &&
				precedence(operators[len(operators)-1]) >= precedence(char) {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, char)
			i++
		} else if char == '(' {
			// Открывающая скобка
			operators = append(operators, char)
			i++
		} else if char == ')' {
			// Закрывающая скобка
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 || operators[len(operators)-1] != '(' {
				return nil, ErrInvalidExpression
			}
			operators = operators[:len(operators)-1]
			i++
		} else {
			return nil, ErrInvalidExpression
		}
	}

	// Добавляем оставшиеся операторы в выходную очередь
	for len(operators) > 0 {
		if operators[len(operators)-1] == '(' || operators[len(operators)-1] == ')' {
			return nil, ErrInvalidExpression
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

// проверяет, является ли символ оператором
func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

// возвращает приоритет оператора
func precedence(char rune) int {
	switch char {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

// вычисление значения выражения в обратной польской нотации
func evaluateRPN(rpn []interface{}) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		switch token := token.(type) {
		case float64:
			stack = append(stack, token)
		case rune:
			if len(stack) < 2 {
				return 0, ErrInvalidExpression
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case '+':
				result = a + b
			case '-':
				result = a - b
			case '*':
				result = a * b
			case '/':
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			default:
				return 0, ErrInvalidExpression
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, ErrInvalidExpression
	}

	return stack[0], nil
}

func main() {
	var str string
	fmt.Scanln(&str)
	res, err := Calc(str)
	fmt.Println(res, err)

}

// // проверка работы функции Calc
// func TestCalc(t *testing.T) {
// 	// Тестовые случаи
// 	testCases := []struct {
// 		expression string
// 		expected   float64
// 		err        error
// 	}{
// 		{"2+3", 5, nil},
// 		{"2+3*4", 14, nil},
// 		{"(2+3)*4", 20, nil},
// 		{"2+3*(4-1)", 11, nil},
// 		{"2/0", 0, errors.New("division by zero")},
// 		{"2+3*(4-1", 0, ErrInvalidExpression},
// 		{"2+3*(4-1))", 0, ErrInvalidExpression},
// 		{"2+3*", 0, ErrInvalidExpression},
// 		{"2+3*(4-1)/0", 0, errors.New("division by zero")},
// 	}

// 	// Проверка каждого тестового случая
// 	for _, tc := range testCases {
// 		result, err := Calc(tc.expression)
// 		if (err == nil && tc.err != nil) || (err != nil && tc.err == nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) || result != tc.expected {
// 			t.Errorf("Calc(%q) = (%f, %v); expected (%f, %v)", tc.expression, result, err, tc.expected, tc.err)
// 		}
// 	}
// }
