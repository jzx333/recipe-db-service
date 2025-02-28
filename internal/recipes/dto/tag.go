package dto

import "time"

type NewTag struct {
	Name  string `db:"name" json:"name"`
	Emoji string `db:"emoji" json:"emoji"`
}

type Tag struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Emoji     string    `db:"emoji" json:"emoji"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
