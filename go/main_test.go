package main

import "testing"

func TestGreetings(t *testing.T) {
	result := greetings("Ola")
	expected := "<b>Ola</b>"
	if result != expected {
			t.Errorf("add() test returned an unexpected result: got %v want %v", result, expected)
	}
}