package env

import (
  "os"
)

import (
  "github.com/bww/go-alert"
)

const (
  defaultIP = "127.0.0.1"
)

/**
 * Immutable variables
 */
var (
  environ string
)

/**
 * Setup immutable environment variables
 */
func init() {
  if v := os.Getenv("ENVIRON"); v != "" {
    environ = v
  }else{
    environ = "devel"
  }
}

/**
 * Determine the current environment
 */
func Environ() string {
  return environ
}

/**
 * Determine our hostname
 */
func Hostname() string {
  var name string
  var err error
  
  if e := Environ(); e != "" && e != "devel" {
    name, err = awsLocalHostname()
    if err != nil {
      alt.Warnf("env: Could not fetch instance hostname from environment: %v", err)
    }else{
      return name
    }
  }
  
  name, err = os.Hostname()
  if err != nil {
    alt.Warnf("Could not obtain hostname from system: %v", err)
  }else{
    return name
  }
  
  return "unknown"
}

/**
 * Determine our local address
 */
func LocalAddr() string {
  addr, err := awsLocalIPv4()
  if err != nil {
    alt.Warnf("env: Could not fetch instance local IPv4 from environment: %v", err)
    addr = defaultIP // punt
  }
  return addr
}

/**
 * Determine our public address
 */
func PublicAddr() string {
  addr, err := awsPublicIPv4()
  if err != nil {
    alt.Warnf("env: Could not fetch instance public IPv4 from environment: %v", err)
    addr = defaultIP // punt
  }
  return addr
}
