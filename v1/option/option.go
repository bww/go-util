package option

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrMalformed = errors.New("Malformed")
	ErrInvalid   = errors.New("Invalid")
)

type Inclusion int

const (
	Include Inclusion = iota
	Require
	Exclude
)

func ParseInclusion(s string) (Inclusion, error) {
	switch strings.ToLower(s) {
	case "include":
		return Include, nil
	case "require":
		return Require, nil
	case "exclude":
		return Exclude, nil
	default:
		return -1, ErrInvalid
	}
}

func (v Inclusion) String() string {
	switch v {
	case Require:
		return "require"
	case Exclude:
		return "exclude"
	case Include:
		return "include"
	default:
		return "invalid"
	}
}

func Included[T any](v T) Optional[T] {
	return Optional[T]{Include, v}
}

func Required[T any](v T) Optional[T] {
	return Optional[T]{Require, v}
}

func Excluded[T any](v T) Optional[T] {
	return Optional[T]{Exclude, v}
}

type Optional[T any] struct {
	Option Inclusion
	Value  T
}

func Parse[T any](p string, f func(string) (T, error)) ([]Optional[T], error) {
	l := parseList(p)
	var o []Optional[T]
	for _, e := range l {
		if len(e) < 1 {
			return nil, ErrMalformed
		}
		var t string
		var n Inclusion
		switch c := e[0]; c {
		case '+':
			n, t = Require, strings.TrimSpace(e[1:])
		case '-':
			n, t = Exclude, strings.TrimSpace(e[1:])
		default:
			if c == '~' {
				n, t = Include, strings.TrimSpace(e[1:])
			} else if c != '+' && c != '-' { // as long as it's not another special character, we consider it the value
				n, t = Include, strings.TrimSpace(e)
			} else {
				return nil, ErrMalformed
			}
		}
		v, err := f(t)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalid, err)
		}
		o = append(o, Optional[T]{n, v})
	}
	return o, nil
}

func parseList(p string) []string {
	var l []string
	if p != "" {
		for _, v := range strings.Split(p, ",") {
			l = append(l, strings.TrimSpace(v))
		}
	}
	return l
}

func (o Optional[T]) String() string {
	switch o.Option {
	case Require:
		return "+" + fmt.Sprint(o.Value)
	case Exclude:
		return "-" + fmt.Sprint(o.Value)
	case Include:
		return "~" + fmt.Sprint(o.Value)
	default:
		return "?" + fmt.Sprint(o.Value)
	}
}
