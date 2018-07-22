package model

import (
	"encoding/json"
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

// GetNoteList 获取笔记列表
func GetNoteList(param CommonParam) (
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
			'id', id,
			'detail', detail
		) AS data,
		COUNT(*) OVER () AS total
		FROM t_note

		ORDER BY id DESC

		LIMIT $1
		OFFSET $2
		`,
		param.Size,
		param.Offset,
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
