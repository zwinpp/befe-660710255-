package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"slices"
)

// Student struct
type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
//var ประกาศแบบ global variable
var students = []Student{
	{ID: "1", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.50},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 3.75},
}

func getStudents(c *gin.Context) {
	yearQuery := c.Query("year")

	if yearQuery != "" {
		fillter := []Student{}

		for _, student := range students {
			//แปลง int to string
			if fmt.Sprint(student.Year) == yearQuery {
				fillter = append(fillter, student)
			}
		}
		c.JSON(http.StatusOK, fillter)
		return
	}
	c.JSON(http.StatusOK, students)
}

func getStudent(c *gin.Context)  {
	id := c.Param("id") // := นิยามตัวแปรใหม่

	for _,student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
}

func creatStudent(c *gin.Context)  {
	var newStudent Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if newStudent.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	if newStudent.Year < 1 || newStudent.Year > 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "year must be 1-4"})
		return
	}
	newStudent.ID = fmt.Sprintf("%d", len(students)+1)

	students = append(students, newStudent)
	c.JSON(http.StatusOK, students)
}

func updateStudent(c *gin.Context)  {
	id := c.Param("id")
	var updateStudent Student

	if err := c.ShouldBindJSON(&updateStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, student := range students {
		if student.ID == id {
			updateStudent.ID = id
			students[i] = updateStudent
			c.JSON(http.StatusOK, updateStudent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
}

func deleteStudent(c *gin.Context)  {
	id := c.Param("id")

	// วน loop student all
	for i, student := range students {
		if student.ID == id {
			students = slices.Delete(students, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "student delete successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"}) // แปลง slice ให้เป็น json
	})

	api := r.Group("/api/v1")
	{
		api.GET("/students", getStudents) //get all
		api.GET("/students/:id", getStudent)
		api.POST("/students", creatStudent)
		api.PUT("/students/:id", updateStudent) // เส้น update
		api.DELETE("/students/:id", deleteStudent)
	}

	r.Run(":8080")
}
