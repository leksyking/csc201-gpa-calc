package main

import (
	"fmt"
)

func main() {
	var students = 3
	for i := 0; i < students; i++ {
		var name string
		fmt.Println("Enter your name: ")
		fmt.Scan(&name)
		moreCourses := "n"
		courses := ""
		grade := 0
		gradeSum := 0
		totalUnit := 0
		for courses != moreCourses {
			var courseName string
			fmt.Println("Enter your course title: ")
			fmt.Scan(&courseName)

			var courseUnit int
			fmt.Println("Enter your course unit: ")
			fmt.Scan(&courseUnit)

			var courseScore int
			fmt.Println("Enter your course score: ")
			fmt.Scan(&courseScore)

			switch {
			case courseScore >= 70:
				grade = 5
			case courseScore >= 60:
				grade = 4
			case courseScore >= 50:
				grade = 3
			case courseScore >= 40:
				grade = 2
			default:
				grade = 1
			}
			gradeSum += grade * courseUnit
			totalUnit += courseUnit

			fmt.Println("Do you want to enter another course?(y/n)")
			fmt.Scan(&courses)
		}

		gradePointAverage := gradeSum / totalUnit
		fmt.Printf("Hello %v!, Your GPA is %.2f.\n", name, float64(gradePointAverage))
	}
}
