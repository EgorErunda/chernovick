package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var students []Student

	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		var singleStudent Student
		err := json.Unmarshal(jsonData, &singleStudent)
		if err != nil {
			return nil, fmt.Errorf("ошибка десериализации JSON: %w", err)
		}
		singleStudent.Grade++
		return json.Marshal(singleStudent)
	}
	for i := range students {
		students[i].Grade++
	}
	return json.Marshal(students)
}
