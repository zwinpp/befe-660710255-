package main

import(
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tea struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64  `json:"price"`
	Weight   int `json:"weight"`
}

var teas = []Tea{
	{ID: "1", Name: "Oolong", Price: 900.00, Weight: 50},
	{ID: "2", Name: "Matcha", Price: 890.00, Weight: 30},
	{ID: "3", Name: "Green Tea", Price: 500.00, Weight: 20},
}

func getTeas(c *gin.Context) {
	teaQuery := c.Query("weight")

	if teaQuery != "" {
		fillter := []Tea{}

		for _, tea := range teas {
			//แปลง int to string
			if fmt.Sprint(tea.Weight) == teaQuery {
				fillter = append(fillter, tea)
			}
		}
		c.JSON(http.StatusOK, fillter)
		return
	}
	c.JSON(http.StatusOK, teas)
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/teas", getTeas)
	}

	r.Run(":8080")
}