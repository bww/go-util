package urls

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		Base       string
		Components []interface{}
		Expect     string
	}{
		{"/", []interface{}{"/a/b/c"}, "/a/b/c"},
		{"", []interface{}{"/a/b/c"}, "/a/b/c"},
		{"/", []interface{}{"a/b/c"}, "/a/b/c"},
		{"////", []interface{}{"a/b/c"}, "////a/b/c"},
		{"/", []interface{}{"////a/b/c"}, "/a/b/c"},
		{"////", []interface{}{"////a/b/c"}, "////a/b/c"},
		{"/", []interface{}{123456}, "/123456"},
		{"/", []interface{}{123456, false}, "/123456/false"},
		{"////", []interface{}{"////a/b/c", 123456, false}, "////a/b/c/123456/false"},
		{"https://api.twilio.com/2010-04-01", []interface{}{"Accounts"}, "https://api.twilio.com/2010-04-01/Accounts"},
		{"https://api.twilio.com/2010-04-01/", []interface{}{"Accounts"}, "https://api.twilio.com/2010-04-01/Accounts"},
		{"https://api.twilio.com/2010-04-01/", []interface{}{"/Accounts"}, "https://api.twilio.com/2010-04-01/Accounts"},
		{"https://api.twilio.com/2010-04-01", []interface{}{"/Accounts"}, "https://api.twilio.com/2010-04-01/Accounts"},
		{"https://api.twilio.com/2010-04-01", []interface{}{"/Accounts", "ABC123", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json"},
		{"https://api.twilio.com/2010-04-01", []interface{}{"/Accounts", "/ABC123/", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json"},
		{"https://api.twilio.com/2010-04-01", []interface{}{"/Accounts/", "/ABC123/", "/Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json"},
		{"https://api.twilio.com/2010-04-01", []interface{}{"/Accounts/", "ABC123/", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json"},
		{"https://api.twilio.com/2010-04-01", []interface{}{"Accounts", "ABC123", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json"},
		{"https://api.twilio.com/2010-04-01", []interface{}{}, "https://api.twilio.com/2010-04-01"},
		{"https://api.twilio.com/2010-04-01/", []interface{}{}, "https://api.twilio.com/2010-04-01/"},
	}
	for _, e := range tests {
		r := Join(e.Base, e.Components...)
		fmt.Println("-->", r)
		assert.Equal(t, e.Expect, r)
	}
}

func TestFile(t *testing.T) {
	tests := []struct {
		Path   string
		Expect string
	}{
		{"", "file:///"},
		{"/", "file:///"},
		{"///", "file:///"},
		{"////////a", "file:///a"},
		{"file://a", "file:///file://a"},
		{"file:///a", "file:///file:///a"},
		{"/a/b/c", "file:///a/b/c"},
	}
	for _, e := range tests {
		r := File(e.Path)
		fmt.Println("-->", r)
		assert.Equal(t, e.Expect, r)
	}
}
