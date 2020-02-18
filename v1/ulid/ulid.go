package ulid

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

// the reader may need to be pooled if there is significant contention; not sure
// if this is actually necessary or possible for crypto/rand (as opposed to math/rand)
var entropy = ulid.Monotonic(rand.Reader, 0)

var Zero ULID = ULID{}

type ULID ulid.ULID

func New() ULID {
	return ULID(ulid.MustNew(ulid.Timestamp(time.Now()), entropy))
}

func Parse(s string) (ULID, error) {
	u, err := ulid.Parse(s)
	if err != nil {
		return Zero, err
	} else {
		return ULID(u), nil
	}
}

func (v ULID) IsZero() bool {
	return v == Zero
}

func (v ULID) String() string {
	return ulid.ULID(v).String()
}

func (v ULID) Time() time.Time {
	return ulid.Time(ulid.ULID(v).Time())
}

func (v ULID) MarshalJSON() ([]byte, error) {
	if v == Zero {
		return []byte("null"), nil
	} else {
		return []byte(`"` + v.String() + `"`), nil
	}
}

func (v *ULID) UnmarshalJSON(data []byte) error {
	r := string(data)
	if r == "null" || r == "0" || r == `""` {
		copy(v[:], Zero[:])
		return nil
	}

	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*v, err = Parse(s)
	if err != nil {
		return err
	}
	return nil
}

func (v ULID) Value() (driver.Value, error) {
	if v == Zero {
		return nil, nil
	} else {
		return v.String(), nil
	}
}

func (v *ULID) Scan(src interface{}) error {
	var err error
	switch c := src.(type) {
	case nil:
		*v = Zero
	case []byte:
		*v, err = Parse(string(c))
	case string:
		*v, err = Parse(c)
	default:
		err = fmt.Errorf("Unsupported type: %T", src)
	}
	return err
}
