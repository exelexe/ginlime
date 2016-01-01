package app

import (
	"ginlime/db"
	//"ginlime/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-dbx"
	"log"
)

type yourHobby struct {
	Name    string `db:"pname"`
	Hobbies string `db:"hname"`
}

func Hobbies(c *gin.Context) {
	db := db.InitDb()
	defer db.Close()

	name := c.Param("person_name")

	query := "SELECT [[persons.name]] AS [[pname]], [[hobbies.name]] AS [[hname]] " +
		"FROM {{persons}}, {{hobbies}} " +
		"WHERE [[persons.id]] = [[hobbies.person_id]] AND [[persons.name]] = {:name}"

	var yourHobbies []yourHobby
	q := db.NewQuery(query)
	q.Bind(dbx.Params{"name": name})
	err := q.All(&yourHobbies)
	if err != nil {
		log.Println(err)
	}
	// for debug
	for k := range yourHobbies {
		log.Println("key: ", k, " value: ", yourHobbies[k])
	}

	c.String(200, "your name is, %s!\n", name)
}
