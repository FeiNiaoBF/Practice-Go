package session

import "foorm/log"

func (s *Session) Begin() (err error) {
	log.Info("transaction begin: ")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error("error transaction begin: ", err)
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit!")
	if err = s.tx.Commit(); err != nil {
		log.Error("error transaction commit: ", err)
		return
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction rollback!")
	if err = s.tx.Rollback(); err != nil {
		log.Error("error transaction rollback: ", err)
		return
	}
	return
}
