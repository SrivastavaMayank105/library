package DB

import(
	"github.com/jinzhu/gorm"
	"fmt"
	lb "LIBRARY"

)

var std lb.Student

func DBInitialization(){
	db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&std)
}