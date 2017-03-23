package goweather

import "testing"

func TestFetch(t *testing.T) {
	m := NewManager()
	_, err := m.Fetch(101200105)
	if err != nil {
		t.Fail()
	}
}
