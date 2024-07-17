package foorm

import (
	"database/sql"
	"foorm/dialect"
	"foorm/log"
	"foorm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// New connect to the `source` database of the `driver`
func New(driver, source string) (engine *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Errorf("Ping the %q failed: %v", driver, err)
		return
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %q Not Found", driver)
		return
	}
	engine = &Engine{db: db, dialect: dial}
	log.Info("connect database success")

	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("failed to close database")
	}
	log.Info("close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
