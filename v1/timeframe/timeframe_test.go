package timeframe

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCompareTimeframe(t *testing.T) {
	tests := []struct {
		TF       Timeframe
		Time     time.Time
		Contains bool
	}{
		{
			TF:       Timeframe{},
			Time:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Contains: true,
		},
		{
			TF:       New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), // inclusive on lower bound
			Contains: true,
		},
		{
			TF:       New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), // exclusive on upper bound
			Contains: false,
		},
		{
			TF:       NewSince(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:     time.Time{},
			Contains: false,
		},
		{
			TF:       NewSince(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:     time.Date(2050, 1, 1, 0, 0, 0, 0, time.UTC),
			Contains: true,
		},
		{
			TF:       NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:     time.Time{},
			Contains: true,
		},
		{
			TF:       NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:     time.Date(2050, 1, 1, 0, 0, 0, 0, time.UTC),
			Contains: false,
		},
	}
	for _, e := range tests {
		fmt.Println(">>>", e.TF, "<>", e.Time)
		assert.Equal(t, e.Contains, e.TF.Contains(e.Time))
	}
}

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
