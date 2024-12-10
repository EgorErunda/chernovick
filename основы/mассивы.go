package main

import "fmt"

func arr() {
	//var arrayName [размерМассива]типДанных
	var numbers [5]int
	var fruits [3]string = [3]string{"apple", "banana", "orange"}
	var fruit = [...]string{"pineapple", "mango", "peach"}
	var apple = fruits[0]
	fmt.Println(numbers, fruits, fruit, apple)
}
