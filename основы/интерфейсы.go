package main

// //Синтаксис определения интерфейса
// type MyInterface interface {
//     Method1(arg1 type1, arg2 type2) returnType1
//     Method2(arg1 type1) returnType2
//     ...
// }

//Пример:
// type Person struct {
//     Name string
//     Age  int
// }

// type Human interface {
//     SayHello()
// }

// func (p Person) SayHello() {
//     fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
// }

// var h Human = Person{Name: "John", Age: 25}
// h.SayHello()

//передача интерфейсов
// type Shape interface {
//     Area() float64
// }

// func CalculateArea(s Shape) float64 {
//     return s.Area()
// }

// type Rectangle struct {
//     Width  float64
//     Height float64
// }

// func (r Rectangle) Area() float64 {
//     return r.Width * r.Height
// }

// rect := Rectangle{Width: 10, Height: 5}
// area := CalculateArea(rect)
