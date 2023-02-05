package template

import (
	"bytes"
	"text/template"
)

type Config struct {
	Funcs template.FuncMap
}

func (c Config) WithOptions(opts []Option) Config {
	for _, opt := range opts {
		c = opt(c)
	}
	return c
}

type Option func(Config) Config

func WithFuncs(f template.FuncMap) Option {
	return func(conf Config) Config {
		conf.Funcs = f
		return conf
	}
}

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

func Parse(f string, opts ...Option) (*Template, error) {
	return ParseWithConfig(f, Config{}.WithOptions(opts))
}

func ParseWithConfig(f string, conf Config) (*Template, error) {
	t := template.New("_")
	if conf.Funcs != nil {
		t.Funcs(conf.Funcs)
	}
	t, err := t.Parse(f)
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
