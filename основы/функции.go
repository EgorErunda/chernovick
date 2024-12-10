package main

import (
	"fmt"
)

//разные сигнатуры функции
//func factorial(n int) int
//func findDiscriminant(a, b, c float64)
//func createUser(name string, age int, weight float64)
//func createUser(age int, name, phoneNumber, secondName string, weight float64)
//func findDiscriminant(a, b, c float64) -> float64 <- тип возвращаемого значения

func calculateAreaAndPerimeter(length, width float64) (float64, float64) {
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}

func devide(a, b int) (int, error) {
	if b == 0 {
		// Возвращаем 0 и ошибку, если возникает деление на 0:
		return 0, fmt.Errorf("wrong divider with value %v", b)
	}

	result := a / b

	return result, nil // Возвращаем результат и nil, если ошибки нет
}

// рeкурсия
func fib(n int) int {
	// крайний случай, так как при f(1) и так далее мы не сможем посчитать обе следующих итерации
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(6)) // Вывод: 8
}
