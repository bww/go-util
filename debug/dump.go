package debug

import (
  "os"
  "io"
  "encoding/json"
)

func Dump(v ...interface{}) {
  Dumpf(os.Stdout, v...)
}

func Dumpf(w io.Writer, v ...interface{}) {
  e := json.NewEncoder(w)
  e.SetIndent("", "  ")
  for _, x := range v {
    e.Encode(x)
  }
}
