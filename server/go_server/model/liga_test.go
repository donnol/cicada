package model

import (
	"testing"
)

func TestMakeSchedule(t *testing.T) {
	name := "Liga"
	kind := "T"
	url := "http://www.laliga.es/en/laliga-santander"
	s := NewSchedule(name, kind, url, []string{})
	if len(s.Teams) != 20 {
		t.Fatal("team total is wrong.")
	}
	if len(s.Result) != 38 {
		t.Fatalf("round total is wrong: %d", len(s.Result))
	}
	for i, round := range s.Result {
		if round.Version != i+1 {
			t.Fatalf("wrong version: %d", round.Version)
		}
		if len(round.Match) != 10 {
			t.Fatalf("wrong match total: %v", round)
		}
	}
	round1 := s.Result[0]
	err := s.Refresh()
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Teams) != 20 {
		t.Fatal("team total is wrong.")
	}
	if len(s.Result) != 38 {
		t.Fatalf("round total is wrong: %d", len(s.Result))
	}
	for i, round := range s.Result {
		if round.Version != i+1 {
			t.Fatalf("wrong version: %d", round.Version)
		}
		if len(round.Match) != 10 {
			t.Fatalf("wrong match total: %v", round)
		}
	}
	round2 := s.Result[0]

	t.Log(round1)
	t.Log(round2)
}
