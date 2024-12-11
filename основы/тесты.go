//Стандартные тесты в Golang пишутся с применением пакета testing
func TestSomething(t *testing.T) {
// ...
}

//Использование t.Run для запуска подтестов

func TestConcurrentFunctionality(t *testing.T) {
    t.Run("Increment", testIncrement)
    t.Run("Decrement", testDecrement)
}

func testIncrement(t *testing.T) {
// тело теста инкрементации...
}

func testDecrement(t *testing.T) {
// тело теста декрементации...
}

//запуск тестов
//go test -v
//запуск отдельных тестов
//go test -v -run=TestPrintHello
//Опция для запуска всех тестов в пакете
//go test -v ./...