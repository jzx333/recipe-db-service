package dto

type Tag struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Emoji     string `json:"emoji"`
	CreatedAt string `json:"created_at"`
}
