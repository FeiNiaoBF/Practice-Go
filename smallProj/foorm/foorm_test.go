package foorm

import (
	"testing"
)

func TestNew(t *testing.T) {
	engine, err := New("sqlite3", "foo.db")
	defer engine.Close()
	if err != nil {
		t.Fatal(err)
	}
}
