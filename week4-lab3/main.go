package main

import (
	"errors"
	"fmt"
)

// ถ้าใช้ในไฟล์ใช้ S or s ก็ได้ แต่ถ้าจะเอาไปใช้ไฟล์อื่น S
type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}
// func ของ struct call method
// bool = boolean
func (s *Student) IsHornor() bool {
	return s.GPA >= 3.50
}
// error = ชนิดข้อมูล errors = package
func (s *Student) Validate() error {
	if s.Name == ""{
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 1-4")
	}
	return nil // nill = null = ไม่ error
}

func main() {
	// var st Student = Student{ID:"1", Name:"sira", Email:"rujirakachornch_S@su.ac.th", Year:4, GPA:4.00}

	// var st Student := Student({ID:"1", Name:"sira", Email:"rujirakachornch_S@su.ac.th", Year:4, GPA:4.00})

	// [4]Student = arrays []Student = slice
	students := []Student{
		{ID:"1", Name:"sira", Email:"rujirakachornch_S@su.ac.th", Year:4, GPA:4.00}, 
		{ID:"2", Name:"mumu", Email:"mumu@email.com", Year:3, GPA:2.50},
	}

	newStudent := Student{ID:"3", Name:"truly", Email:"truly@email.com", Year:4, GPA:3.50}
	students = append(students, newStudent)

	// _ ใส่แทน i ถ้าไม่ต้องการใช้ตัวแปรไปใช้
	for i, student:= range students {
		fmt.Printf("%d Hornor = %v\n", i,student.IsHornor())
		fmt.Printf("%d Validation = %v\n", i, student.Validate())
	}
	
}
