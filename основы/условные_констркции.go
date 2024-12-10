package main

import (
	"fmt"
)

func uslovie() {
	var score int
	var number = 15
	fmt.Scanln(&score)
	//long form
	if score >= 60 {
		fmt.Println("Отличная работа!")
	} else {
		fmt.Println("Нужно усерднее учиться!")
	}
	//short form
	if square := number * number; square > 50 {
		fmt.Println("Квадрат числа больше 50")
	} else {
		fmt.Println("Квадрат числа меньше или равен 50")
	}
}

func switch_statement() {
	weather := "rain"
	switch weather {
	case "rain":
		fmt.Println("Погода дождливая")
	case "sunny":
		fmt.Println("Погода солнечная")
	default:
		fmt.Println("Погода не определена")
	}
}
