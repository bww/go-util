package ulid

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeULID(t *testing.T) {
	v := New()
	tests := []struct {
		ULID   *ULID
		Expect string
		Error  error
	}{
		{
			&v, fmt.Sprintf(`"%v"`, v), nil,
		},
		{
			&Zero, `null`, nil,
		},
		{
			nil, `null`, nil,
		},
	}

	for _, e := range tests {
		fmt.Println(">>>", e)
		d, err := json.Marshal(e.ULID)
		if e.Error != nil {
			assert.Equal(t, e.Error, err)
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			assert.Equal(t, e.Expect, string(d))
		}
	}
}

func TestDecodeULID(t *testing.T) {
	v := New()
	tests := []struct {
		Text   string
		Expect ULID
		Error  error
	}{
		{
			fmt.Sprintf(`"%v"`, v), v, nil,
		},
		{
			`""`, Zero, nil,
		},
		{
			`0`, Zero, nil,
		},
		{
			`null`, Zero, nil,
		},
	}

	for _, e := range tests {
		fmt.Println(">>>", e)
		var u ULID
		err := json.Unmarshal([]byte(e.Text), &u)
		if e.Error != nil {
			assert.Equal(t, e.Error, err)
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			assert.Equal(t, e.Expect, u)
		}
	}
}
