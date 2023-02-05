package template

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

var funcs = template.FuncMap{
	"hello": func() string { return "hello!" },
}

func TestTemplateOneshot(t *testing.T) {
	tests := []struct {
		Tmpl   string
		Data   interface{}
		Expect []byte
		Error  bool
	}{
		{
			Tmpl:   "Hello, nobody",
			Data:   nil,
			Expect: []byte("Hello, nobody"),
		},
		{
			Tmpl: "Hello, {{ .Name }}",
			Data: struct {
				Name string
			}{
				"Bobzoo",
			},
			Expect: []byte("Hello, Bobzoo"),
		},
		{
			Tmpl: "Hello, {{ .Missing }}",
			Data: struct {
				Name string
			}{
				"Bobzoo",
			},
			Error: true,
		},
	}
	for _, e := range tests {
		res, err := Exec(e.Tmpl, e.Data)
		if e.Error {
			fmt.Println("***", err)
			assert.NotNil(t, err, fmt.Sprint(err))
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			fmt.Println("-->", string(e.Expect))
			assert.Equal(t, e.Expect, res)
		}
	}
}

func TestTemplateWithFuncs(t *testing.T) {
	tests := []struct {
		Tmpl   string
		Data   interface{}
		Expect []byte
		Error  bool
	}{
		{
			Tmpl:   "Oh, {{ hello }}",
			Expect: []byte("Oh, hello!"),
		},
	}
	for _, e := range tests {
		tmpl, err := Parse(e.Tmpl, WithFuncs(funcs))
		assert.NoError(t, err)
		res, err := tmpl.Exec(e.Data)
		if e.Error {
			fmt.Println("***", err)
			assert.NotNil(t, err, fmt.Sprint(err))
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			fmt.Println("-->", string(e.Expect))
			assert.Equal(t, e.Expect, res)
		}
	}
}
