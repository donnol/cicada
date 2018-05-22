package model

import (
	"time"
)

// Note 笔记
type Note struct {
	ID        int
	UserID    int
	Detail    string
	CreatedAt time.Time
}

// AddNote 添加笔记
func AddNote(note Note) (id int, err error) {
	err = _db.Get(&id, `INSERT INTO t_note(user_id, detail)
		VALUES($1, $2)
		RETURNING id
		`,
		note.UserID,
		note.Detail,
	)
	if err != nil {
		return
	}
	return
}
