package main

import (
	"encoding/json"
	"fmt"
)

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var mergeData []map[string]interface{}

	for _, jsonData := range jsonDataList {
		var data []map[string]interface{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			return nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
		}
		mergeData = append(mergeData, data...)
	}

	mergeJSON, err := json.Marshal(mergeData)
	if err != nil {
		return nil, fmt.Errorf("ошибка сериализации JSON: %w", err)
	}

	return mergeJSON, nil
}
