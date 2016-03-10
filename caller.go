package ui

import (
	"reflect"
)

type Caller struct {
	PubEvent *Eventer
	PriEvent *Eventer
	module   Module
}

func (e *Caller) MakeParam(args ...interface{}) []reflect.Value {
	leng := len(args)
	ret := make([]reflect.Value, leng+1)
	ret[0] = reflect.ValueOf(e.module)
	if leng > 0 {
		for k, v := range args {
			ret[k+1] = reflect.ValueOf(v)
		}
		return ret
	}
	return ret
}

func (c *Caller) Call(n string, args ...interface{}) {
	c.CallByValues(n, c.MakeParam(args...))
}

func (c *Caller) CallByValues(n string, args []reflect.Value) {
	ret := c.PubEvent.Call(n, args)
	if ret {
		c.PriEvent.Call(n, args)
	}
}

func NewCaller(m Module) *Caller {
	c := &Caller{
		module: m,
	}

	c.PubEvent = NewEventer(m)
	c.PriEvent = NewDefaultEventer(m)

	return c
}
