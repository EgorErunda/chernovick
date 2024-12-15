package main

import (
	"encoding/json"
	"fmt"
)

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	//хуевина для хранения сгруппированных данных
	classMap := make(map[string][]map[string]interface{})

	//хуевина для десериализации данных в массив объкетов
	var data []map[string]interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
	}
	for _, item := range data {
		//получаем значение поля class
		class, ok := item["class"].(string)
		if !ok {
			return nil, fmt.Errorf("поле 'class' отсутствует или имеет неверный тип")
		}
		classMap[class] = append(classMap[class], item)
	}
	// Сериализация каждого списка обратно в JSON
	result := make(map[string][]byte)

	for class, items := range classMap {
		jsonData, err := json.Marshal(items)
		if err != nil {
			return nil, fmt.Errorf("ошибка сериализации JSON для класса: %w", err)
		}
		result[class] = jsonData
	}
	return result, nil
}
