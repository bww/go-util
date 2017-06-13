package trace

import (
  "io"
  "os"
  "fmt"
  "time"
  "math"
  "strings"
)

type aggregate func([]time.Duration)(time.Duration)

var displayUnit time.Duration
var groupByName aggregate

func init() {
  switch strings.ToLower(os.Getenv("GOUTIL_TRACE_GROUP_SPANS_BY")) {
    case "none":        // nothing
    case "avg", "mean": groupByName = mean
    case "max":         groupByName = max
    case "sum":         groupByName = sum
    default:            groupByName = sum
  }
  switch strings.ToLower(os.Getenv("GOUTIL_TRACE_DURATION_UNITS")) {
    case  "s":        displayUnit = time.Second
    case "ms":        displayUnit = time.Millisecond
    case "us", "μs":  displayUnit = time.Microsecond
    default:          displayUnit = time.Nanosecond
  }
}

// An individual span
type Span struct {
  Name      string
  Started   time.Time
  Duration  time.Duration
  Aggregate int
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
func New(n string) *Trace {
  return &Trace{Name:n}
}

// Begin a new span
func (t *Trace) Start(n string) *Span {
  s := &Span{n, time.Now(), 0, 0}
  t.Spans = append(t.Spans, s)
  return s
}

// Write a trace to the specified writer
func (t *Trace) Write(w io.Writer) (int, error) {
  
  // group by name, use the position of the first occurrance
  spans := t.Spans
  if groupByName != nil {
    spans = group(groupByName, spans)
  }
  
  // compute the trace duration
  var et, lt time.Time
  var sd time.Duration
  var si int
  for i, e := range spans {
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
    s = fmt.Sprintf("%v (%v in %d spans; longest: #%d @ %s)\n", t.Name, td, len(spans), si + 1, formatDuration(sd))
  }else{
    s = fmt.Sprintf("%v (no closed spans)\n", t.Name)
  }
  
  if l := len(spans); l > 0 {
    nd := int(math.Log10(float64(l + 1))) + 1
    nf := fmt.Sprintf("%%%dd", nd)
    
    var dm int
    var ds []string
    for _, e := range spans {
      var d string
      if e.Duration > 0 {
        d = formatDuration(e.Duration)
      }else {
        d = "(open)"
      }
      ds = append(ds, d)
      if l := len(d); l > dm {
        dm = l
      }
    }
    
    df := fmt.Sprintf("%%%ds", dm)
    for i, e := range spans {
      s += fmt.Sprintf("  #"+ nf +" "+ df +" ", i + 1, ds[i])
      s += e.Name
      if e.Aggregate > 1 {
        s += fmt.Sprintf(" (⨉%d)", e.Aggregate)
      }
      s += "\n"
    }
  }
  
  return fmt.Fprint(w, s)
}

// Group spans using the specified aggregate function
func group(a aggregate, s []*Span) []*Span {
  base := make([]*Span, len(s))
  copy(base, s)
  
  for i := 0; i < len(base); i++ {
    b := base[i]
    m := []time.Duration{b.Duration}
    for j := i + 1; j < len(base); {
      if c := base[j]; c.Name == b.Name {
        m = append(m, c.Duration)
        for k := j + 1; k < len(base); k++ { base[k-1] = base[k] }
        base = base[:len(base)-1]
      }else{
        j++
      }
    }
    if len(m) > 1 {
      base[i] = &Span{Name:b.Name, Started:b.Started, Duration:a(m), Aggregate:len(m)}
    }
  }
  
  return base
}

// Format a duration
func formatDuration(d time.Duration) string {
  if displayUnit == time.Nanosecond {
    return d.String()
  }else{
    return fmt.Sprintf("%f", float64(d) / float64(displayUnit)) + unitSuffix(displayUnit)
  }
}

// Obtain the unit suffix
func unitSuffix(u time.Duration) string {
  switch u {
    case time.Second: return "s"
    case time.Millisecond: return "ms"
    case time.Microsecond: return "μs"
    case time.Nanosecond: return "ns"
    default: return "?s"
  }
}

// Mean aggregate function
func mean(d []time.Duration) time.Duration {
  var t time.Duration
  for _, e := range d {
    t += e
  }
  return t / time.Duration(len(d))
}

// Max aggregate function
func max(d []time.Duration) time.Duration {
  var m time.Duration
  for _, e := range d {
    if e > m {
      m = e
    }
  }
  return m
}

// Sum aggregate function
func sum(d []time.Duration) time.Duration {
  var s time.Duration
  for _, e := range d {
    s += e
  }
  return s
}

