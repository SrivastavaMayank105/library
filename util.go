package LIBRARY

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var std []Student

type Student struct{
	gorm.Model
	NAME string `json:"name" binding:"required"`
	ID int `json:"id"  binding:"required"`
}

func (s *Student)  CreateStudent(c *gin.Context) {
	var p Student
		c.BindJSON(&p)
		std=append(std,p)
}

func (s *Student) DisplayAll() []Student{
	return std
}

func (s *Student) GetStudent(id int) Student {
	for _,v:= range std {
		if id == v.ID{
			return v
		}
	}
return Student{}
}