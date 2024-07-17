package dialect

import (
	"reflect"
	"sync"
)

// Dialect 抽象出各个数据库差异的部分
type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []any)
}

var dialects = map[string]Dialect{}
var mu sync.Mutex

func RegisterDialect(name string, dialect Dialect) {
	mu.Lock()
	dialects[name] = dialect
	mu.Unlock()
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	mu.Lock()
	defer mu.Unlock()
	dialect, ok = dialects[name]

	return
}
