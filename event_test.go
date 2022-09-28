package event

import "testing"

func TestNewEvent(t *testing.T) {
	handlers := []Condition{Compare, Compare}
	args := []Argument{
		{20, GreaterThanEqual, 10},
		{10, GreaterThanEqual, 10},
	}
	b, err := NewEvent(1, 1).Do(handlers, args, LogicAnd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b, err)
}
