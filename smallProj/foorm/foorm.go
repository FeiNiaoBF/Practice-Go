package foorm

import (
	"database/sql"
	"fmt"
	"foorm/dialect"
	"foorm/log"
	"foorm/session"
	"strings"
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

type TxFunc func(*session.Session) (any, error)

// Transaction use transaction
func (engine *Engine) Transaction(f TxFunc) (result any, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			_ = s.Rollback() // err is non-nil; don't change it
		} else {
			defer func() {
				if err != nil {
					_ = s.Rollback()
				}
			}()
			err = s.Commit()
			// err is nil; if Commit returns error update err
		}
	}()

	return f(s)
}

// Migrate 变更已有的表
func (engine *Engine) Migrate(value any) error {
	_, err := engine.Transaction(func(s *session.Session) (result any, err error) {
		if !s.Model(value).HasTable() {
			log.Infof("[Migrate]: table %s doesn't exist", s.RefTable().Name)
			return nil, s.CreateTable()
		}
		table := s.RefTable()
		rows, _ := s.Raw(fmt.Sprintf("SELECT * FROM %s LIMIT 1", table.Name)).QueryRows()
		cols, _ := rows.Columns()
		addCols := difference(table.FieldNames, cols)
		delCols := difference(cols, table.FieldNames)
		log.Infof("[Migrate]: added cols %v, deleted cols %v", addCols, delCols)

		for _, col := range addCols {
			f := table.GetField(col)
			builderStr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", table.Name, f.Name, f.Type)
			if _, err = s.Raw(builderStr).Exec(); err != nil {
				return
			}
		}

		if len(delCols) == 0 {
			return
		}
		tmp := "tmp_" + table.Name
		fieldStr := strings.Join(table.FieldNames, ", ")
		s.Raw(fmt.Sprintf("CREATE TABLE %s AS SELECT %s from %s;", tmp, fieldStr, table.Name))
		s.Raw(fmt.Sprintf("DROP TABLE %s;", table.Name))
		s.Raw(fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", tmp, table.Name))
		_, err = s.Exec()
		return
	})
	return err
}

// difference returns a - b
func difference(a []string, b []string) (diff []string) {
	mapB := make(map[string]bool)
	for _, v := range b {
		mapB[v] = true
	}
	for _, v := range a {
		if _, ok := mapB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}
