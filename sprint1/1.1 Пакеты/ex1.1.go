package main

import (
	"errors"
	"fmt"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("размеры игрового поля должны быть положительными числами")
	}
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	world := &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}

	return world, nil
}

func main() {
	// Пример использования функции NewWorld
	world, err := NewWorld(-1, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Создано игровое поле: Height=%d, Width=%d\n", world.Height, world.Width)
}
