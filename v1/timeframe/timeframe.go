package timeframe

import (
	"errors"
	"strings"
	"time"
)

const (
	encodedLayout = "2006-01-02T15:04:05.999999999Z07:00" // this is: time.RFC3339Nano, but guaranteed not to change
	sep           = ".."
)

var errMalformed = errors.New("Malformed")

// A timeframe with no bounds, representing all of time
var Forever = Timeframe{}

// A timeframe
type Timeframe struct {
	Since *time.Time `json:"since" db:"since"`
	Until *time.Time `json:"until" db:"until"`
}

// Create a timeframe
func New(f, t time.Time) Timeframe {
	return Timeframe{&f, &t}
}

// Create a timeframe from -inf to the specified time
func NewUntil(t time.Time) Timeframe {
	return Timeframe{nil, &t}
}

// Create a timeframe from the specified time to +inf
func NewSince(f time.Time) Timeframe {
	return Timeframe{&f, nil}
}

// Parse a timeframe from its encoded format
func Parse(s string) (Timeframe, error) {
	if s == "" { // empty string is a valid unbounded timeframe
		return Timeframe{}, nil
	}

	var l, r string
	if x := strings.Index(s, sep); x < 0 {
		return Timeframe{}, errMalformed
	} else {
		l, r = s[:x], s[x+len(sep):]
	}

	var since *time.Time
	if l != "" {
		v, err := time.Parse(encodedLayout, l)
		if err != nil {
			return Timeframe{}, err
		}
		since = &v
	}

	var until *time.Time
	if r != "" {
		v, err := time.Parse(encodedLayout, r)
		if err != nil {
			return Timeframe{}, err
		}
		until = &v
	}

	return Timeframe{
		Since: since,
		Until: until,
	}, nil
}

// Is a timeframe finite (is bounded)
func (t Timeframe) IsFinite() bool {
	return t.Since != nil || t.Until != nil
}

// What is the duration of the timeframe. If a timeframe is not bounded on both ends, the
// duration is infinite, which is expressed as a negative value. If Until is before Since,
// the value will also be negative, this is a logically invalid state that is interpreted
// as the timeframe being unbounded.
func (t Timeframe) Duration() time.Duration {
	if !t.IsFinite() {
		return -1
	} else {
		return t.Until.Sub(*t.Since)
	}
}

// Determine if the timeframe contains the specified time. Timeframes are inclusive on the
// lower bound and exclusive on the upper bound.
func (t Timeframe) Contains(a time.Time) bool {
	if t.Since != nil && a.Sub(*t.Since) < 0 {
		return false
	}
	if t.Until != nil && a.Sub(*t.Until) >= 0 {
		return false
	}
	return true
}

func (t Timeframe) String() string {
	var s string
	if t.Since != nil {
		s += (*t.Since).Format(encodedLayout)
	}
	s += sep
	if t.Until != nil {
		s += (*t.Until).Format(encodedLayout)
	}
	return s
}

func (t Timeframe) Format(layout string) string {
	var s string
	if t.Since != nil {
		s += (*t.Since).Format(layout)
	} else {
		s += "the beginning of time"
	}
	s += " until "
	if t.Until != nil {
		s += (*t.Until).Format(layout)
	} else {
		s += "the end of time"
	}
	return s
}

// CompareDuration compares the duration between t.Since and t.Until with the given duration
func (t Timeframe) CompareDuration(d time.Duration) int {
	if t.Until == nil || t.Since == nil {
		return 1
	}
	to, from := *(t.Until), *(t.Since)
	if v := to.Sub(from); v > d {
		return 1
	} else if v < d {
		return -1
	}
	return 0
}
