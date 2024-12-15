package main

import (
	"fmt"
	"reflect"
)

// Vehicle Интерфейс для транспортного средства
type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

// Car Структура для автомобиля
type Car struct {
	Speed    float64 // Скорость автомобиля в км/ч
	Type     string  // Тип транспортного средства
	FuelType string  // Тип топлива
}

// CalculateTravelTime Реализация метода для расчёта времени в пути для автомобиля
func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

// Motorcycle Структура для мотоцикла
type Motorcycle struct {
	Speed float64 // Скорость мотоцикла в км/ч
	Type  string  // Тип транспортного средства
}

// CalculateTravelTime Реализация метода для расчёта времени в пути для мотоцикла
func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

// EstimateTravelTime Функция для оценки времени в пути для списка транспортных средств
func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	results := make(map[string]float64)
	for _, vehicle := range vehicles {
		// Используем полиморфизм: вызываем метод CalculateTravelTime для каждого транспортного средства
		travelTime := vehicle.CalculateTravelTime(distance)
		// Добавляем результат в карту с ключом, соответствующим типу транспортного средства
		results[reflect.TypeOf(vehicle).String()] = travelTime
	}
	return results
}

func main() {
	// Создаём список транспортных средств
	vehicles := []Vehicle{
		Car{Speed: 60, Type: "Седан", FuelType: "Бензин"},
		Motorcycle{Speed: 80, Type: "Мотоцикл"},
	}

	// Расстояние до пункта назначения
	distance := 200.0 // в километрах

	// Оцениваем время в пути для каждого транспортного средства
	travelTimes := EstimateTravelTime(vehicles, distance)

	// Выводим результаты
	for vehicleType, time := range travelTimes {
		fmt.Printf("Время в пути для %s: %.2f часов\n", vehicleType, time)
	}
}
