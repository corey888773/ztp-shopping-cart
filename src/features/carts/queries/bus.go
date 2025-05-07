package queries

import (
	"errors"
	"reflect"
)

type Handler interface {
	Handle(query interface{}) (interface{}, error)
}

type QueryBus struct {
	handlers map[reflect.Type]Handler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[reflect.Type]Handler),
	}
}

func (qb *QueryBus) Register(query interface{}, handler Handler) {
	t := reflect.TypeOf(query)
	qb.handlers[t] = handler
}

func (qb *QueryBus) Handle(query interface{}) (interface{}, error) {
	queryType := reflect.TypeOf(query)
	if handler, ok := qb.handlers[queryType]; ok {
		return handler.Handle(query)
	}
	return nil, errors.New(ErrInvalidQuery)
}
