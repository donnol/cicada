package model

import (
	"encoding/json"
	"time"
)

// Note 笔记
type Note struct {
	ID        int
	UserID    int `db:"user_id"`
	Title     string
	Detail    string
	CreatedAt time.Time `db:"created_at"`
}

// AddNote 添加笔记
func AddNote(note Note) (id int, err error) {
	err = _db.Get(&id, `INSERT INTO t_note(user_id, title, detail)
		VALUES($1, $2, $3)
		RETURNING id
		`,
		note.UserID,
		note.Title,
		note.Detail,
	)
	if err != nil {
		return
	}
	return
}

// ModifyNote 修改笔记
func ModifyNote(note Note) (err error) {
	_, err = _db.Exec(`Update t_note set
		title = $1,
		detail = $2
		Where id = $3
		`,
		note.Title,
		note.Detail,
		note.ID,
	)
	if err != nil {
		return
	}
	return
}

// GetNoteList 获取笔记列表
func GetNoteList(note Note, param CommonParam) (
	res struct {
		Data  []json.RawMessage
		Total int
	},
	err error,
) {
	var dbResult []struct {
		Data  json.RawMessage
		Total int
	}
	err = _db.Select(&dbResult, `
		SELECT json_build_object(
			'ID', id,
			'Title', title,
			'Detail', detail,
			'CreatedAt', created_at
		) AS data,
		COUNT(*) OVER () AS total
		FROM t_note

		WHERE true
		
		AND CASE WHEN $3 <> '' THEN
			title ~* $3
		ELSE true END

		AND CASE WHEN $4 <> 0 THEN
			id = $4
		ELSE true END

		ORDER BY id DESC

		LIMIT $1
		OFFSET $2
		`,
		param.Size,
		param.Offset,
		note.Title,
		note.ID,
	)
	if err != nil {
		return
	}
	for i, single := range dbResult {
		if i == 0 {
			res.Total = single.Total
		}
		res.Data = append(res.Data, single.Data)
	}
	return
}

// GetNote 获取笔记
func GetNote(id int) (note Note, err error) {
	err = _db.Get(&note, `
		SELECT id, user_id, title, detail, created_at
		FROM t_note
		WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return
	}
	return
}
