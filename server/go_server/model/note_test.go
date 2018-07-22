package model

import (
	"encoding/json"
	"testing"
)

func TestGetNoteList(t *testing.T) {
	for _, cas := range []CommonParam{
		{Size: 10, Offset: 0},
		{Size: 1, Offset: 0},
		{Size: 1, Offset: 1},
	} {
		res, err := GetNoteList(cas)
		if err != nil {
			t.Fatal(err)
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("res is %s\n", resJSON)
	}
}
