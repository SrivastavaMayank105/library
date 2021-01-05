package Startup

import (

	"github.com/gin-gonic/gin"
	pg "LIBRARY/Pages"
)
func Startup(){
	r:=gin.Default()
	r.POST("/signup",pg.Signup)
	r.GET("/students",pg.Users)
	r.GET("/student/:id",pg.StudentID)
	r.Run(":8080")
}
