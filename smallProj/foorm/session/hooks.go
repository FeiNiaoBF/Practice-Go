package session

import (
	"foorm/log"
	"reflect"
)

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// hooks interface

type BeforeQueryInterface interface {
	BeforeQuery(s *Session) error
}
type AfterQueryInterface interface {
	AfterQuery(s *Session) error
}
type BeforeUpdateInterface interface {
	BeforeUpdate(s *Session) error
}
type AfterUpdateInterface interface {
	AfterUpdate(s *Session) error
}
type BeforeDeleteInterface interface {
	BeforeDelete(s *Session) error
}
type AfterDeleteInterface interface {
	AfterDelete(s *Session) error
}
type BeforeInsertInterface interface {
	BeforeInsert(s *Session) error
}
type AfterInsertInterface interface {
	AfterInsert(s *Session) error
}

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value any) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}
