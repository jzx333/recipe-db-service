package dto

import "time"

type Tag struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Emoji     string    `db:"emoji" `
	CreatedAt time.Time `db:"created_at"`
}

type NewTag struct {
	Name  string `db:"name" json:"name"`
	Emoji string `db:"emoji" json:"emoji"`
}
