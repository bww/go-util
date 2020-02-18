package rand

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test env parsing
func TestRandomStringFromset(t *testing.T) {
	sets := []string{
		"abcd1234",
		"_G_4_h_z",
		Alpha,
		Uppercase,
		Lowercase,
		Digit,
		AlphaNumeric,
		Punctuation,
		Password,
	}
	for _, set := range sets {
		for i := 0; i < 10; i++ {
			r := RandomStringFromSet(64, set)
			fmt.Println(">>>", r)
			for _, e := range r {
				assert.Equal(t, true, strings.IndexRune(set, e) >= 0)
			}
		}
	}
}
