package foorm

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestNew(t *testing.T) {
	engine, err := New("sqlite3", "foo.db")
	defer engine.Close()
	if err != nil {
		t.Fatal(err)
	}
}
