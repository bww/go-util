package option

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOptional(t *testing.T) {
	tests := []struct {
		Text   string
		Expect []Optional[int]
		Error  error
	}{
		{
			Text:  "/",
			Error: ErrInvalid,
		},
		{
			Text:  "not a number",
			Error: ErrInvalid,
		},
		{
			Text:   "",
			Expect: []Optional[int](nil),
		},
		{
			Text:   "+100",
			Expect: []Optional[int]{{Option: Require, Value: 100}},
		},
		{
			Text:   "-100",
			Expect: []Optional[int]{{Option: Exclude, Value: 100}},
		},
		{
			Text:   "~100",
			Expect: []Optional[int]{{Option: Include, Value: 100}},
		},
		{
			Text:   "100",
			Expect: []Optional[int]{{Option: Include, Value: 100}},
		},
	}

	for _, e := range tests {
		m, err := Parse(e.Text, strconv.Atoi)
		if e.Error != nil {
			assert.True(t, errors.Is(err, e.Error), err)
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			assert.Equal(t, e.Expect, m, e.Text)
		}
	}
}
