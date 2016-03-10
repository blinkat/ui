package ui

import (
	"fmt"
	"reflect"
)

// win list
var modules = make(map[Handle]Module)

// register window for call event
func Register(m Module) {
	modules[m.Handle()] = m
}

// event const
const (
	_EVENT_CLOSE   = "Close"
	_EVENT_PAINT   = "Paint"
	_EVENT_CREATED = "Created"
	_EVENT_SHOW    = "Show"
)

func NewEventer(m Module) *Eventer {
	e := &Eventer{
		module: m,
	}
	e._type = reflect.ValueOf(e).Elem()
	return e
}

// normal event type
// return true call default event
// return false do nothing
type Event func(Module) bool
type PaintEvent func(Module, *DeviceContext) bool

// event
type Eventer struct {
	_type  reflect.Value
	module Module

	Close   Event
	Paint   PaintEvent
	Created Event
	Show    Event
}

// **********************************
func (e *Eventer) Call(n string, args []reflect.Value) bool {
	met := e._type.FieldByName(n)
	ret := true

	for _, v := range args {
		if v.Kind() == reflect.Invalid {
			fmt.Println(v.Type(), "is valid")
		}
	}

	if !met.IsNil() {
		r := met.Call(args)
		ret = r[0].Bool()
	}
	return ret
}

// ***********************
func callModuleEvent(h Handle, n string, args ...interface{}) {
	if v, ok := modules[h]; ok {
		v.Events().Call(n, args...)
	}
}

func NewDefaultEventer(m Module) *Eventer {
	e := NewEventer(m)

	e.Close = func(md Module) bool {
		md.Destory()
		return false
	}

	return e
}
