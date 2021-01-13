package template

import (
	"bytes"
	"text/template"
)

func Exec(f string, d interface{}) ([]byte, error) {
	t, err := Parse(f)
	if err != nil {
		return nil, err
	}
	return t.Exec(d)
}

type Template struct {
	*template.Template
}

func Parse(f string) (*Template, error) {
	t, err := template.New("_").Parse(f)
	if err != nil {
		return nil, err
	}
	return &Template{t}, nil
}

func (t *Template) Exec(d interface{}) ([]byte, error) {
	b := &bytes.Buffer{}
	err := t.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
