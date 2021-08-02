package db

import "testing"

func TestSave(t *testing.T) {
	rep := NewInMemoryDB(nil)
	if err := rep.Save("key1", 1); err != nil {
		t.Error("Must save without errors")
	}

	_, err := rep.Find("key1")

	if err == ErrRecordNotFound {
		t.Error("Must find a non-nil element")
	}
}
