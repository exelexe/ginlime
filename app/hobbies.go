package app

import (
	"ginlime/db"
	"ginlime/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-dbx"
	"log"
)

func Hobbies(c *gin.Context) {
	db := db.InitDb()
	defer db.Close()

	name := c.Param("person_name")

	var person structs.Person
	q := db.NewQuery("SELECT [[id]], [[name]] FROM {{persons}} WHERE [[name]]={:name}")
	q.Bind(dbx.Params{"name": name})
	err := q.One(&person)
	if err != nil {
		log.Println(err)
	}

	//log.Println("person: ", person.Name, person.Id)

	c.String(200, "your name is, %s!\n", name)
}
