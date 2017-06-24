package env

import (
  "os"
  "fmt"
  "strings"
  "io/ioutil"
)

// Load the contents of an .env file into the current environment
func LoadEnvfile(f string) error {
  e, err := ReadEnvfile(f)
  if err != nil {
    return err
  }
  for k, v := range e {
    err = os.Setenv(k, v)
    if err != nil {
      return err
    }
  }
  return nil
}

// Read the contents of an .env file
func ReadEnvfile(f string) (map[string]string, error) {
  
  r, err := os.Open(f)
  if err != nil {
    return nil, err
  }
  
  d, err := ioutil.ReadAll(r)
  if err != nil {
    return nil, err
  }
  
  e := make(map[string]string)
  s := string(d)
  
  p := 0
  for i := 0; i <= len(s); i++ {
    if i == len(s) || s[i] == '\n' {
      k, v, err := envDecl(s[p:i])
      if err != nil {
        return nil, err
      }
      e[k] = os.ExpandEnv(v)
    }
  }
  
  return e, nil
}

// Read an environment declaration, in the form:
//   KEY=VALUE # Maybe
func envDecl(s string) (string, string, error) {
  var err error
  
  if strings.Contains(s, "#") {
    qc, nx, nq := byte(0), 0, 0
    for i := 0; i < len(s); i++ {
      switch s[i] {
        case '\\':
          nx++; continue
        case qc:
          if nx % 2 == 0 {
            nq++
          }
        case '\'', '"':
          if qc == 0 {
            qc = s[i]; nq++
          }
        case  '#':
          if nq % 2 == 0 {
            s = s[:i]
            break
          }
      }
    }
  }
  
  x := strings.Index(s, "=")
  if x < 0 {
    return "", "", fmt.Errorf("Invalid decl: %v", s)
  }
  
  k, v := strings.TrimSpace(s[:x]), strings.TrimSpace(s[x+1:])
  if len(v) > 0 && (v[0] == '\'' || v[0] == '"') {
    v, err = unquote(v)
    if err != nil {
      return "", "", err
    }
  }
  
  return k, v, nil
}

// Unquote a string
func unquote(s string) (string, error) {
  var u string
  if len(s) < 2 {
    return u, nil
  }
  
  q := s[0]
  if s[len(s) - 1] != q {
    return "", fmt.Errorf("Unbalanced quote: %s", s)
  }
  
  ex := false
  for i := 1; i < len(s); i++ {
    switch s[i] {
      case '\\':
        if ex { u += "\\" }
        ex = !ex
      case q:
        if ex {
          u += string(q)
          ex = false
        }else if i != len(s) - 1 {
          return "", fmt.Errorf("Invalid quote: %s", s[:i])
        }
      default:
        u += string(s[i])
    }
  }
  
  return u, nil
}
