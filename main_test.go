package main

import (
	"myapp/doctor"
	"testing"
)

func TestMain(t *testing.T) {
	got := doctor.Intro()
	want := "I'm Eliza ---------Talk to the program by typing in plain English, using normal upper and lower-case letters and punctuation.  Enter 'quit' when done. Hello. How are you feeling today?"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
