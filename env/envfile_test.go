package env

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

// Test env decl parsing
func TestEnvDecl(t *testing.T) {
  assertEnvDecl(t, "KEY=VAL", "KEY", "VAL")
  assertEnvDecl(t, "KEY=VAL # Blah...", "KEY", "VAL")
  assertEnvDecl(t, "KEY=\"VAL\" # Blah...", "KEY", "VAL")
  assertEnvDecl(t, "KEY=\"VAL\"\n# Blah...", "KEY", "VAL")
  assertEnvDecl(t, "KEY=\"WHY VAL\"\n# Blah...", "KEY", "WHY VAL")
  assertEnvDecl(t, "KEY=\"WHY!VAL\"\n# Blah...", "KEY", "WHY!VAL")
  assertEnvDecl(t, "KEY='VAL' # Blah...", "KEY", "VAL")
  assertEnvDecl(t, "KEY='VAL'\n# Blah...", "KEY", "VAL")
  assertEnvDecl(t, "KEY='WHY VAL'\n# Blah...", "KEY", "WHY VAL")
  assertEnvDecl(t, "KEY='WHY!VAL'\n# Blah...", "KEY", "WHY!VAL")
  
  assertEnvDecl(t, `KEY="#" # Blah...`, "KEY", "#")
  assertEnvDecl(t, `KEY="\"#\"\"" # Blah...`, "KEY", `"#""`)
  assertEnvDecl(t, `KEY='"#""' # Blah...`, "KEY", `"#""`)
  assertEnvDecl(t, `KEY="'#'" # Blah...`, "KEY", `'#'`)
  assertEnvDecl(t, `KEY="\\#" # Blah...`, "KEY", `\#`)
}

// Assert env decl
func assertEnvDecl(t *testing.T, s, ek, ev string) {
  fmt.Print(s, " -> ")
  ak, av, err := envDecl(s)
  if assert.Nil(t, err, fmt.Sprintf("%v", err)) {
    fmt.Println(ak, av)
    assert.Equal(t, ek, ak)
    assert.Equal(t, ev, av)
  }
}
