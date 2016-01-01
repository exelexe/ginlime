package structs

import (
	"time"
)

type Hobby struct {
	Id        int64     `db:"id"`
  PersonId  int64     `db:"person_id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
