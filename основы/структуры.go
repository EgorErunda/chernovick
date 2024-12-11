package main

//Объявление структуры
type Student struct {
    name string
    age  int
}

student1 := Student{name: "vasya", age: 15}

//метод структуры
type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Scale(factor float64) (float64, float64) {
    r.width *= factor
    r.height *= factor
    return r.width, r.height
}