package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Courses struct{
	ID string `json:"id"`
	Title string `json:"name"`
}

var courses = []Courses{
	{"100","Grokking Modern System Design "},
	{"101","CloudLab: WebSockets-based Chat Application using API Gateway"},
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func getSpecificCourse(c *gin.Context){
	id := c.Param("id")
	for _, course := range courses{
		if course.ID == id{
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Курс с таким ID не найден"})
}

func addCourse(c *gin.Context){
	var courseToAdd Courses
	err := c.ShouldBindJSON(&courseToAdd)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	courses = append(courses, courseToAdd)
	c.IndentedJSON(http.StatusCreated, courses)
}

func main() {
	r := gin.Default()
	r.GET("/courses", getCourses)
	r.GET("/courses/:id", getSpecificCourse)
	r.POST("/courses", addCourse)

	r.Run(":8080")
}
