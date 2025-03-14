package notes

import (
	"database/sql"
	"fmt"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) RGetNoteOwner(id int) (string, error) {
	var UserId Note
	err := r.db.QueryRow("SELECT user_id FROM notes WHERE id=$1", id).Scan(&UserId.UserId)
	if err != nil {
		return "", err
	}
	return UserId.UserId, nil
}

func (r *Repo) RGetNotes(id int) ([]Note, error) {
	var notes []Note
	rows, err := r.db.Query(`SELECT id, title, body, created_at, updated_at FROM notes WHERE user_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.Id, &n.Title, &n.Body, &n.Created, &n.Updated); err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}
	if len(notes) == 0 {
		return nil, fmt.Errorf("user`s notes not found")
	}
	return notes, nil
}

func (r *Repo) RGetNoteById(id int) (*Note, error) {
	var note Note
	err := r.db.QueryRow(`SELECT id, title, body, created_at, updated_at FROM notes WHERE id=$1`, id).
		Scan(&note.Id, &note.Title, &note.Body, &note.Created, &note.Updated)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *Repo) RCreateNewNote(title, body string, id int) error {
	_, err := r.db.Exec(`INSERT INTO notes (title, body, user_id, created_at) VALUES ($1,$2,$3,$4)`, title, body, id, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) RDeleteNote(id int) error {
	_, err := r.db.Exec(`DELETE FROM notes WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) RUpdateNote(title, body string, id int) error {
	_, err := r.db.Exec(`UPDATE notes SET title=$1, body=$2, updated_at=$4 WHERE id=$3`, title, body, id, time.Now())
	if err != nil {
		return err
	}
	return nil
}
