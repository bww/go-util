package text

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

// Assert
func assertRanges(t *testing.T, s string, r [][]int, e error) {
  v, err := ParseRanges(s, []int{0,100})
  if e != nil {
    assert.Equal(t, e, err)
  }else if assert.Nil(t, err, fmt.Errorf("%v", err)) {
    assert.Equal(t, r, v)
  }
}

// Test ranges
func TestRanges(t *testing.T) {
  assertRanges(t, "2-4", [][]int{{2,4}}, nil)
  assertRanges(t, "2-4,5-6", [][]int{{2,4},{5,6}}, nil)
  assertRanges(t, "2-4,4-6", [][]int{{2,6}}, nil)
  assertRanges(t, "2-4,4,5-6", [][]int{{2,6}}, nil)
  assertRanges(t, "2 - 4, 4, 5 - 6", [][]int{{2,6}}, nil)
  assertRanges(t, "-", [][]int{{0,100}}, nil)
  assertRanges(t, "10-", [][]int{{10,100}}, nil)
  assertRanges(t, "-20", [][]int{{0,20}}, nil)
  assertRanges(t, "-4,4,5-", [][]int{{0,100}}, nil)
  assertRanges(t, "", nil, ErrEmptyInput)
}
