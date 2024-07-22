package foorm

import (
	"errors"
	"foorm/session"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	engine, err := New("sqlite3", "foo.db")
	defer engine.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func openDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := New("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

type User struct {
	Name string `foorm:"PRIMARY KEY"`
	Age  int
}

func transactionRollback(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return nil, errors.New("Error")
	})
	// 最终事务回滚，表创建失败
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		// ...
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}

func TestEngine_Migrate(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text PRIMARY KEY, XXX integer);").Exec()
	_, _ = s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	engine.Migrate(&User{})

	rows, _ := s.Raw("SELECT * FROM User").QueryRows()
	columns, _ := rows.Columns()
	if !reflect.DeepEqual(columns, []string{"Name", "Age"}) {
		t.Fatal("Failed to migrate table User, got columns", columns)
	}
}
