package session

import (
	"foorm/log"
	"testing"
)

type Account struct {
	ID       int `foorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before inert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSessionSpy().Model(&Account{})

	_ = s.DropTable()
	_ = s.CreateTable()

	row, _ := s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})
	t.Logf("row: %v", row)

	u := &Account{}
	want := &Account{1001, "******"}
	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatalf("Failed to call hooks after query, got %v, but want %v", u, want)
	}
}
