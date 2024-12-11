//Пример:
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
result, err := divide(10, 0)
if err != nil {
	fmt.Println("Ошибка:", err)
	return
}
fmt.Println("Результат:", result)

//В Go есть стандартный пакет errors,
//который позволяет создавать новые ошибки
//с помощью функции New()
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

//добнее выносить ошибки как константы
var (
    ErrDivisionByZero = errors.New("division by zero")
)

//Defer(Отложенный вызов)
func createFile(filename string, data []byte) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(data)
    if err != nil {
        return err
    }

    return nil
}

