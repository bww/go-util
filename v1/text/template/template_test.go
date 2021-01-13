package template

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateOneshot(t *testing.T) {
	tests := []struct {
		Tmpl   string
		Data   interface{}
		Expect []byte
		Error  bool
	}{
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
