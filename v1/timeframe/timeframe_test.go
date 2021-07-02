package timeframe

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEncodeTimeframe(t *testing.T) {
	tests := []struct {
		TF  Timeframe
		Enc string
	}{
		{
			TF:  Timeframe{},
			Enc: "..",
		},
		{
			TF:  New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			Enc: "2020-01-01T00:00:00Z..2021-01-01T00:00:00Z",
		},
		{
			TF:  New(time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 1, time.UTC)),
			Enc: "2020-01-01T00:00:00.000000001Z..2021-01-01T00:00:00.000000001Z",
		},
		{
			TF:  NewSince(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Enc: "2020-01-01T00:00:00Z..",
		},
		{
			TF:  NewSince(time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC)),
			Enc: "2020-01-01T00:00:00.000000001Z..",
		},
		{
			TF:  NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Enc: "..2020-01-01T00:00:00Z",
		},
		{
			TF:  NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC)),
			Enc: "..2020-01-01T00:00:00.000000001Z",
		},
	}
	for _, e := range tests {
		enc := e.TF.String()
		fmt.Println(">>>", e.TF, "→", enc)
		if assert.Equal(t, e.Enc, enc) {
			dec, err := Parse(enc)
			if assert.Nil(t, err, fmt.Sprint(err)) {
				fmt.Println("<<<", enc, "→", dec)
				assert.Equal(t, e.TF, dec)
			}
		}
	}
}
