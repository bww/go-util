package scan
// This package is convenient, but it is not designed to be especially
// performant. In particular, strings are concatenated naively rather
// than using buffers. For short input it should be suitable for most
// uses, however.

import (
  "fmt"
  "unicode"
  "unicode/utf8"
)

var (
  ErrInvalidSequence = fmt.Errorf("Invalid sequence")
)

// Parse a parameter string
func String(p string, q, x rune) (string, string, error) {
  var s string
  var err error
  
  if c, w := utf8.DecodeRuneInString(p); c != q {
    return "", p, ErrInvalidSequence
  }else{
    p = p[w:]
  }
  
outer:
  for len(p) > 0 {
    c, w := utf8.DecodeRuneInString(p)
    p = p[w:]
    switch c {
      case q:
        break outer
      case x:
        var r rune
        r, p, err = unescape(p, q, x)
        if err != nil {
          return "", p, err
        }
        s += string(r)
      default:
        s += string(c)
    }
  }
  
  return s, p, nil
}

// Unescape escape sequences in a string
func unescape(p string, q, x rune) (rune, string, error) {
  switch c, w := utf8.DecodeRuneInString(p); c {
    case 'a':
      return '\a', p[w:], nil
    case 'b':
      return '\b', p[w:], nil
    case 'f':
      return '\f', p[w:], nil
    case 'n':
      return '\n', p[w:], nil
    case 'r':
      return '\r', p[w:], nil
    case 't':
      return '\t', p[w:], nil
    case 'v':
      return '\v', p[w:], nil
    case x, q:
      return c, p[w:], nil
    default:
      return 0, p, fmt.Errorf("Unsupported escape sequence: \\%v in '%v'", string(c), p)
  }
}

// Escape escapable chars in a string
func escape(s string, q, x rune) string {
  var o string
  for _, e := range s {
    switch e {
      case '\a': o += string(x) +"a"
      case '\b': o += string(x) +"b"
      case '\f': o += string(x) +"f"
      case '\n': o += string(x) +"n"
      case '\r': o += string(x) +"r"
      case '\t': o += string(x) +"t"
      case '\v': o += string(x) +"v"
      case    q: o += string(x) + string(q)
      default:   o += string(e)
    }
  }
  return o
}

// Skip past leading whitespace
func skipWhite(s string) string {
  var i int
  var e rune
  for i, e = range s {
    if !unicode.IsSpace(e) {
      break
    }
  }
  return s[i:]
}
