package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMaybe(v int, err error) (int, error) {
	return v, err
}

func TestMust(t *testing.T) {
	assert.Equal(t, 123, Must(testMaybe(123, nil)))
}
