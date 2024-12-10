package main

import (
	"fmt"
	"sort"
)

//Как устроен слайс изнутри? Это структура с тремя полями:
// type slice struct {
// 	array unsafe.Pointer - указатель на элемент массива, с которого начинается слайс
// 	len   int            - длина
// 	cap   int            - ёмкость
// }

func slice() {
	//Создание слайса
	a := make([]int, 2) // создаём слайс целых чисел длиной 2
	a[0] = 0
	a[1] = 1
	fmt.Println(a) // [0 1]

	//Объявление переменной через литерал или var
	var c []int           // объявляем переменную типа «слайс целых чисел»
	fmt.Println(c == nil) // выведет true

	b := []int{0, 1, 1, 2} // присваиваем переменной b новый слайс
	fmt.Println(b)         // [0 1 1 2]

	//Срез уже существующего слайса или массива
	x := [5]int{0, 1, 1, 2, 3} // создаём массив из 5 целых чисел
	y := a[:4]                 // берём слайс из 4 элементов от массива
	z := b[:3]                 // берём слайс из 3 элементов от слайса

	fmt.Println(x) // [0 1 1 2 3]
	fmt.Println(y) // [0 1 1 2]
	fmt.Println(z) // [0 1 1]

	//Длина и ёмкость слайса(length, capacity)
	f := []int{0, 1, 1, 2, 3}
	fmt.Println(len(f), cap(f)) // 5, 5

	q := f[:3]                  // [0, 1, 1]
	fmt.Println(len(q), cap(q)) // 3, 5

	//append(slice, value)

	//Сортировка
	intSlice := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	fmt.Println(sort.IntsAreSorted(intSlice)) // Проверим, отсортирован ли слайс
	sort.Ints(intSlice)                       // Сама сортировка
	fmt.Println(intSlice)                     // [1 2 3 4 5 6 7 8 9]

	stringSlice := []string{"Go", "Bravo", "Gopher", "YaLyceum", "Alpha", "Grin", "Delta"}
	sort.Strings(stringSlice) // Сортировка уже для строк
	fmt.Println(stringSlice)  // [Alpha Bravo Delta Go Gopher Grin YaLyceum]
}
