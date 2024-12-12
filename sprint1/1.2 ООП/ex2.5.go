package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	reportID := rand.Intn(200000)
	return Report{
		User:     user,
		ReportID: reportID,
		Date:     reportDate,
	}
}

func PrintReport(report Report) {
	fmt.Printf("Отчет ID: %d\n", report.ReportID)
	fmt.Printf("Пользователь: %s (ID-%d)\n", report.Name, report.ID)
	fmt.Printf("Email: %s\n", report.Email)
	fmt.Printf("Возраст: %d\n", report.Age)
	fmt.Printf("Дата создания отчета: %s\n", report.Date)
	fmt.Println("************************************************")
}

func GenerateUserReports(users []User, reportDate string) []Report {
	var reports []Report
	for _, user := range users {
		report := CreateReport(user, reportDate)
		reports = append(reports, report)
	}
	return reports
}
