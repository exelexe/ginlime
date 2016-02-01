package app

import (
	"encoding/json"
	"ginlime/db"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-dbx"
	"log"
)

type yourHobby struct {
	Name  string `db:"pname"`
	Hobby string `db:"hname"`
	Level string `db:"level"`
}

type responseJson struct {
	Hobbies  []hobby `json:"hobbies"`
	YourName string  `json:"your_name"`
}

type hobby struct {
	Name string `json:"name"`
	Level string `json:"level"`
}

func Hobbies(c *gin.Context) {
	// get db data
	data := selectHobby(c.Param("person_name"))
	// output response
	c.String(200, genJson(data))
}

func selectHobby(name string) []yourHobby {
	// open a database
	db := db.InitDbx()
	defer db.Close()

	var yhs []yourHobby
	q := db.NewQuery(
		"SELECT [[persons.name]] AS [[pname]], [[hobbies.name]] AS [[hname]], [[hobbies.level]] " +
		"FROM {{persons}}, {{hobbies}} " +
		"WHERE [[persons.id]] = [[hobbies.person_id]] AND [[persons.name]] = {:name}")
	q.Bind(dbx.Params{"name": name})
	err := q.All(&yhs)
	if err != nil {
		log.Println(err)
	}

	return yhs
}

func genJson(data []yourHobby) string {
	var hobbies []hobby
	for i := range data {
		var h hobby
		h.Name = data[i].Hobby
		h.Level = data[i].Level
		hobbies = append(hobbies, h)
	}

	r := responseJson{}
	r.YourName = data[0].Name
	r.Hobbies = hobbies

	jstr, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		log.Println(err)
	}

	return string(jstr)
}
