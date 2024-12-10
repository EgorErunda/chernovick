package main

import (
	"fmt"
)

func cycle_for() {
	for i := 0; i < 10; i++ {
		fmt.Println("Привет")
	}
}

func cycle_while() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

//break continue

func Range() {
	for i := range 10 {
		fmt.Println(i)
	}

	for i, letter := range "Hello, world!" {
		fmt.Println(i, string(letter))
	}
}
