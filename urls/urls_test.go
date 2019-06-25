package urls

import (
	"testing"
)

func TestJoin(t *testing.T) {
	checkJoin(t, "/", []interface{}{"/a/b/c"}, "/a/b/c")
	checkJoin(t, "", []interface{}{"/a/b/c"}, "/a/b/c")
	checkJoin(t, "/", []interface{}{"a/b/c"}, "/a/b/c")
	checkJoin(t, "////", []interface{}{"a/b/c"}, "////a/b/c")
	checkJoin(t, "/", []interface{}{"////a/b/c"}, "/a/b/c")
	checkJoin(t, "////", []interface{}{"////a/b/c"}, "////a/b/c")
	checkJoin(t, "/", []interface{}{123456}, "/123456")
	checkJoin(t, "/", []interface{}{123456, false}, "/123456/false")
	checkJoin(t, "////", []interface{}{"////a/b/c", 123456, false}, "////a/b/c/123456/false")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{"Accounts"}, "https://api.twilio.com/2010-04-01/Accounts")
	checkJoin(t, "https://api.twilio.com/2010-04-01/", []interface{}{"Accounts"}, "https://api.twilio.com/2010-04-01/Accounts")
	checkJoin(t, "https://api.twilio.com/2010-04-01/", []interface{}{"/Accounts"}, "https://api.twilio.com/2010-04-01/Accounts")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{"/Accounts"}, "https://api.twilio.com/2010-04-01/Accounts")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{"/Accounts", "ABC123", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{"/Accounts", "/ABC123/", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{"/Accounts/", "/ABC123/", "/Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{"/Accounts/", "ABC123/", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{"Accounts", "ABC123", "Messages.json"}, "https://api.twilio.com/2010-04-01/Accounts/ABC123/Messages.json")
	checkJoin(t, "https://api.twilio.com/2010-04-01", []interface{}{}, "https://api.twilio.com/2010-04-01")
	checkJoin(t, "https://api.twilio.com/2010-04-01/", []interface{}{}, "https://api.twilio.com/2010-04-01/")
}

func checkJoin(t *testing.T, a string, b []interface{}, e string) {
	if r := Join(a, b...); r != e {
		t.Errorf("Invalid endpoint; for: [%v; %v], expected: %v, got: %v", a, b, e, r)
	}
}
