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

func TestCompareOrdering(t *testing.T) {
	tests := []struct {
		TF     Timeframe
		Time   time.Time
		Expect Ordering
	}{
		{
			TF:     Timeframe{},
			Time:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: Within,
		},

		{
			TF:     NewSince(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: Within,
		},
		{
			TF:     NewSince(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: Within,
		},
		{
			TF:     NewSince(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: Before,
		},

		{
			TF:     NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:   time.Date(2010, 12, 31, 23, 59, 59, 999, time.UTC),
			Expect: Within,
		},
		{
			TF:     NewUntil(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: Within,
		},
		{
			TF:     NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Time:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: After,
		},

		{
			TF: New(
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			Time:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: Within,
		},
		{
			TF: New(
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			Time:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: Before,
		},
		{
			TF: New(
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			Time:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Expect: After,
		},
	}
	for i, e := range tests {
		assert.Equal(t, e.Expect, e.TF.Compare(e.Time), "#%d: %v <> %v", i, e.TF, e.Time)
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

func TestMarshalTimeframeColumn(t *testing.T) {
	tests := []struct {
		TF  Timeframe
		Enc []byte
	}{
		{
			TF:  Timeframe{},
			Enc: []byte(".."),
		},
		{
			TF:  New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			Enc: []byte("2020-01-01T00:00:00Z..2021-01-01T00:00:00Z"),
		},
		{
			TF:  New(time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC), time.Date(2021, 1, 1, 0, 0, 0, 1, time.UTC)),
			Enc: []byte("2020-01-01T00:00:00.000000001Z..2021-01-01T00:00:00.000000001Z"),
		},
		{
			TF:  NewSince(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Enc: []byte("2020-01-01T00:00:00Z.."),
		},
		{
			TF:  NewSince(time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC)),
			Enc: []byte("2020-01-01T00:00:00.000000001Z.."),
		},
		{
			TF:  NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			Enc: []byte("..2020-01-01T00:00:00Z"),
		},
		{
			TF:  NewUntil(time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC)),
			Enc: []byte("..2020-01-01T00:00:00.000000001Z"),
		},
	}
	for _, e := range tests {
		enc, err := e.TF.MarshalColumn()
		if assert.NoError(t, err) {
			fmt.Println(">>>", e.TF, "→", string(enc))
			if assert.Equal(t, e.Enc, enc) {
				var dec Timeframe
				err = dec.UnmarshalColumn(enc)
				if assert.NoError(t, err) {
					fmt.Println("<<<", string(enc), "→", dec)
					assert.Equal(t, e.TF, dec)
				}
			}
		}
	}
}
