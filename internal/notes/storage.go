package notes

import "time"

type Note struct {
	Id      int       `json:"id"`
	UserId  string    `json:"user_id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}
