package trace

import (
  "io"
  "fmt"
  "time"
  "math"
)

// An individual span
type Span struct {
  Name      string
  Started   time.Time
  Duration  time.Duration
}

// Finish a span
func (s *Span) Finish() {
  s.Duration = time.Since(s.Started)
}

// A trace, which manages a set of related spans
type Trace struct {
  Name  string
  Spans []*Span
}

// Create a trace
func NewTrace(n string) *Trace {
  return &Trace{Name:n}
}

// Begin a new span
func (t *Trace) Start(n string) *Span {
  s := &Span{n, time.Now(), 0}
  t.Spans = append(t.Spans, s)
  return s
}

// Write a trace to the specified writer
func (t *Trace) Write(w io.Writer) (int, error) {
  
  var et, lt time.Time
  var sd time.Duration
  var si int
  for i, e := range t.Spans {
    if i == 0 || e.Started.Before(et) {
      et = e.Started
    }
    if a := e.Started.Add(e.Duration); a.After(lt) {
      lt = a
    }
    if e.Duration > sd {
      sd = e.Duration
      si = i
    }
  }
  
  var s string
  if td := lt.Sub(et); td > 0 {
    s = fmt.Sprintf("%v (%v in %d spans; longest: %d @ %v)\n", t.Name, td, len(t.Spans), si + 1, sd)
  }else{
    s = fmt.Sprintf("%v (no closed spans)\n", t.Name)
  }
  
  if l := len(t.Spans); l > 0 {
    nd := int(math.Log10(float64(l + 1))) + 1
    nf := fmt.Sprintf("%%%dd", nd)
    
    var dm int
    var ds []string
    for _, e := range t.Spans {
      var d string
      if e.Duration > 0 {
        d = fmt.Sprintf("%v", e.Duration)
      }else {
        d = "(open)"
      }
      ds = append(ds, d)
      if l := len(d); l > dm {
        dm = l
      }
    }
    
    df := fmt.Sprintf("%%%ds", dm)
    for i, e := range t.Spans {
      s += fmt.Sprintf("  #"+ nf +" "+ df +" %v", i + 1, ds[i], e.Name)
      s += "\n"
    }
  }
  
  return fmt.Fprint(w, s)
}
