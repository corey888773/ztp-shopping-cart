package commands

import (
	"errors"
	"fmt"
	"reflect"
)

type Handler interface {
	Handle(cmd interface{}) error
}

type CommandBus struct {
	handlers map[reflect.Type]Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[reflect.Type]Handler),
	}
}

func (cb *CommandBus) Register(cmd interface{}, handler Handler) {
	t := reflect.TypeOf(cmd)
	cb.handlers[t] = handler
}

func (cb *CommandBus) Handle(command interface{}) error {
	cmdType := reflect.TypeOf(command)
	if handler, ok := cb.handlers[cmdType]; ok {
		return handler.Handle(command)
	}
	// print cmdType and map of handlers
	fmt.Printf("Command type: %v\n", cmdType)
	fmt.Printf("Registered handlers: %v\n", cb.handlers)

	return errors.New(ErrInvalidCommand)
}
