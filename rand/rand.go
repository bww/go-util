package rand

import (
	"crypto/rand"
	"fmt"
	"io"
	"net"
	"strings"
)

// Common sets for random strings
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alpha        = Uppercase + Lowercase
	Digit        = "0123456789"
	AlphaNumeric = Alpha + Digit
	Punctuation  = "!@#$%^&*~?=-+_./|"
	Password     = AlphaNumeric + Punctuation
)

// Cached MAC address
var macaddr net.HardwareAddr

func init() {
	// search our network interfaces for a hardware MAC address
	if interfaces, err := net.Interfaces(); err == nil {
		for _, i := range interfaces {
			if (i.Flags&net.FlagLoopback) == 0 && len(i.HardwareAddr) > 0 {
				macaddr = i.HardwareAddr
				break
			}
		}
	}
	// if we failed to obtain the MAC address of the current computer, we will
	// use a randomly generated 6 byte sequence instead and set the multicast
	// bit as recommended in RFC 4122.
	if macaddr == nil {
		macaddr = make(net.HardwareAddr, 6)
		randomBytes(macaddr)
		macaddr[0] = macaddr[0] | 0x01
	}
}

// Obtain the current host's MAC address
func HardwareAddr() net.HardwareAddr {
	return macaddr
}

// Obtain the current host's MAC address as a hex string
func HardwareKey() string {
	return fmt.Sprintf("%x", macaddr)
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	randomBytes(b)
	return b
}

func ReadRandom(b []byte) {
	randomBytes(b)
}

func randomBytes(b []byte) {
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err.Error()) // rand should never fail
	}
}

func RandomString(n int) string {
	return RandomStringFromSet(n, AlphaNumeric)
}

func RandomStringFromSet(n int, set string) string {
	b := RandomBytes(n)
	l := len(set)
	s := &strings.Builder{}
	for i := 0; i < len(b); i++ {
		s.WriteRune(rune(set[int(b[i])%l]))
	}
	return s.String()
}
