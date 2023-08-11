package exec_test

import (
	"testing"

	"github.com/MarioJim/gonja/exec"
	"github.com/stretchr/testify/assert"
)

var ctxCases = []struct {
	name     string
	value    any
	asString string
	flags    flags
}{
	{"nil", nil, "", flags{IsNil: true}},
	{"string", "Hello World", "Hello World", flags{IsString: true, IsTrue: true}},
	{"int", 42, "42", flags{IsInteger: true, IsNumber: true, IsTrue: true}},
	{"int 0", 0, "0", flags{IsInteger: true, IsNumber: true}},
	{"float", 42., "42.000000", flags{IsFloat: true, IsNumber: true, IsTrue: true}},
	{"float 0.0", 0., "0.000000", flags{IsFloat: true, IsNumber: true}},
	{"true", true, "True", flags{IsBool: true, IsTrue: true}},
	{"false", false, "False", flags{IsBool: true}},
}

func TestContext(t *testing.T) {
	for _, cc := range ctxCases {
		test := cc
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Error(err)
				}
			}()
			assert := assert.New(t)

			ctx := exec.EmptyContext()
			ctx.Set(test.name, test.value)
			value, ok := ctx.Get(test.name)

			assert.Equal(test.value, value)
			assert.True(ok)
		})
	}
}

func TestSubContext(t *testing.T) {
	for _, cc := range ctxCases {
		test := cc
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Error(err)
				}
			}()
			assert := assert.New(t)

			ctx := exec.EmptyContext()
			ctx.Set(test.name, test.value)
			sub := ctx.Inherit()
			value, ok := sub.Get(test.name)

			assert.Equal(test.value, value)
			assert.True(ok)
		})
	}
}

func TestFuncContext(t *testing.T) {
	ctx := exec.EmptyContext()
	ctx.Set("func", func() {})

	cases := []struct {
		name string
		ctx  *exec.Context
	}{
		{"top context", ctx},
		{"sub context", ctx.Inherit()},
	}

	for _, c := range cases {
		test := c
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Error(err)
				}
			}()
			assert := assert.New(t)

			value, ok := test.ctx.Get("func")
			assert.True(ok)
			_, ok = value.(func())
			assert.True(ok)
		})
	}
}
