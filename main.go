package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Course struct{
	ID string `json:"id"`
	Title string `json:"name"`
}

var courses = []Course{
	{"100","Grokking Modern System Design "},
	{"101","CloudLab: WebSockets-based Chat Application using API Gateway"},
}

func updateCourse(c *gin.Context){
	courseIdtoUpdate := c.Param("id")
	var updatedCourse Course
	c.BindJSON(&updatedCourse)

	for index, course := range courses {
		if course.ID == courseIdtoUpdate {
			courses[index] = updatedCourse
			c.IndentedJSON(http.StatusOK, courses)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found!"})
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
	var courseToAdd Course
	err := c.ShouldBindJSON(&courseToAdd)
	if err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	courses = append(courses, courseToAdd)
	c.IndentedJSON(http.StatusCreated, courses)
}

func deleteCourse(c *gin.Context){
	var id = c.Param("id")
	for index, course := range courses{
		if course.ID == id{
			courses = append(courses[:index], courses[index + 1:]...)
			c.IndentedJSON(http.StatusOK, courses)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found!"})
}

func main() {
	r := gin.Default()
	r.GET("/courses", getCourses)
	r.GET("/courses/:id", getSpecificCourse)
	r.POST("/courses", addCourse)
	r.PUT("/course/:id", updateCourse)
	r.DELETE("course/:id", deleteCourse)

	r.Run(":8080")
}
