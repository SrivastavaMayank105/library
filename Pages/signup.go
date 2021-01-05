package Pages

import (
	"strconv"
	"github.com/gin-gonic/gin"
	lb "LIBRARY"
)

var (
	std lb.Student
)

func Signup(c *gin.Context){
	c.Header("Content-Type","application/json")	
	//c.BindJSON(&std)
	std.CreateStudent(c)
		c.JSON(200,gin.H{
			"message":"Signup successfull",
		})
	
}

func Users(c *gin.Context){
	c.Header("Content-type","application/json")
	s1:=std.DisplayAll()
	c.JSON(200,gin.H{
		"users":s1,
	})
}


func StudentID(c *gin.Context){
	c.Header("Content-type","application/json")
	query:=c.Param("id")
	q,_:=strconv.Atoi(query)
	d:=std.GetStudent(q)
	if d.NAME!="" {
		c.JSON(200,gin.H{
			"user name": d.NAME,
			"user id": d.ID,
		})
		return 
	}
	c.JSON(404,gin.H{
		"message":"User not found",
	})
}