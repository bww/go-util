package text

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Test indent string
 */
func TestIndent(t *testing.T) {
	assert.Equal(t, "> Hello", Indent("Hello", "> "))
	assert.Equal(t, "> Hello\n> There", Indent("Hello\nThere", "> "))
	assert.Equal(t, "> Hello\n> There\n> ", Indent("Hello\nThere\n", "> "))
	assert.Equal(t, "> Hello, let’s include some → unicode codepoints!\n> There\n> ", Indent("Hello, let’s include some → unicode codepoints!\nThere\n", "> "))
}

/**
 * Test indent writer
 */
func TestIndentWriter(t *testing.T) {
	b := &bytes.Buffer{}
	w := NewIndentWriter("> ", IndentOptionIndentFirstLine, b)
	io.WriteString(w, "Hello\nThere\nBr")
	io.WriteString(w, "ah\nChillin.\n")
	io.WriteString(w, "Hello, let’s include some → unicode")
	io.WriteString(w, " codepoints!\nThere\n")
	assert.Equal(t, "> Hello\n> There\n> Brah\n> Chillin.\n> Hello, let’s include some → unicode codepoints!\n> There\n> ", string(b.Bytes()))
}
