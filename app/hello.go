package app

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := c.Param("name")
	c.String(200, "hello-3, %s!\n", name)
}

// exsample Gorp
//func Hobbies(c *gin.Context) {
//  // initialize the DbMap
//	dbmap := db.InitGorp()
//	defer dbmap.Db.Close()
//
//  // get HTTP Request params
//	name := c.Param("person_name")
//
//	var persons []structs.Person
//	_, err := dbmap.Select(&persons, "SELECT id FROM persons WHERE name=?", name)
//	if err != nil {
//		log.Println(err)
//	}
//
//	for x, p := range persons {
//	        log.Println("    %d: %v\n", x, p)
//	}
//	//log.Println("person: ", params.Name)
//
//	c.String(200, "your name is, %s!\n", name)
//}
