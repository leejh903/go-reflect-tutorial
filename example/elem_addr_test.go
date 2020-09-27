package example

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Alpha string
	Beta  int
}

func (Foo) bar() {}

type Bar interface {
	bar()
}

var foo = Foo{
	Alpha: "alpha",
	Beta:  2,
}

func Test_Elem_Pointer(t *testing.T) {
	// Elem() returns origin value of interface or pointer
	// it panics if origin value is not a pointer
	val := reflect.ValueOf(foo)
	assert.Panics(t, func() { val.Elem() })

	// so it should be a pointer, not struct
	val = reflect.ValueOf(&foo)
	assert.NotPanics(t, func() { val.Elem() })
}

func Test_Elem_Interface(t *testing.T) {
	var bar Bar = foo
	val := reflect.ValueOf(bar)
	assert.Panics(t, func() { val.Elem() })

	// pointing struct of interface should be a pointer to call Elem()
	bar = &foo
	val = reflect.ValueOf(bar)
	assert.NotPanics(t, func() { val.Elem() })
}

// Addr() is like &foo in code
func Test_Addr(t *testing.T) {
	// `foo` is not a pointer
	val := reflect.ValueOf(foo)
	assert.False(t, val.CanAddr())
	assert.Panics(t, func() { val.Addr() })

	// when we call reflect.ValueOf(s), it is a copy of `s`
	// so it panics as well
	val = reflect.ValueOf(&foo)
	assert.False(t, val.CanAddr())
	assert.Panics(t, func() { val.Addr() })

	// when we call Elem(), we get origin value of `foo`
	val = val.Elem()
	assert.True(t, val.CanAddr())
	assert.NotPanics(t, func() { val.Addr() })
}
