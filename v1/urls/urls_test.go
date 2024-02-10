package urls

import (
	"fmt"
	"net/url"
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

func TestMergeQuery(t *testing.T) {
	tests := []struct {
		Base   string
		Query  []url.Values
		Expect string
		Error  error
	}{
		{
			"https://api.twilio.com/2010-04-01/",
			[]url.Values{
				{"a": []string{"b"}},
			},
			"https://api.twilio.com/2010-04-01/?a=b",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?a=b",
			[]url.Values{},
			"https://api.twilio.com/2010-04-01/?a=b",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?a=b",
			nil,
			"https://api.twilio.com/2010-04-01/?a=b",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?",
			[]url.Values{
				{"a": []string{"b"}},
			},
			"https://api.twilio.com/2010-04-01/?a=b",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?a=b",
			[]url.Values{
				{"a": []string{"b"}},
			},
			"https://api.twilio.com/2010-04-01/?a=b&a=b",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?a=b",
			[]url.Values{
				{"a": []string{"b", "c", "d"}},
			},
			"https://api.twilio.com/2010-04-01/?a=b&a=b&a=c&a=d",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?a=b",
			[]url.Values{
				{"a": []string{"c"}},
				{"a": []string{"d"}},
			},
			"https://api.twilio.com/2010-04-01/?a=b&a=c&a=d",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?a=b",
			[]url.Values{
				{"a": []string{"b"}},
				{"a": []string{"c"}},
			},
			"https://api.twilio.com/2010-04-01/?a=b&a=b&a=c",
			nil,
		},
		{
			"https://api.twilio.com/2010-04-01/?a=b",
			[]url.Values{
				{"a": []string{"b"}},
				{"a": []string{"c"}},
				{"a": []string{"c"}},
			},
			"https://api.twilio.com/2010-04-01/?a=b&a=b&a=c&a=c",
			nil,
		},
	}
	for _, e := range tests {
		r, err := MergeQuery(e.Base, e.Query...)
		if e.Error != nil {
			fmt.Println("***", err)
			assert.Equal(t, e.Error, err)
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			fmt.Println("-->", r)
			assert.Equal(t, e.Expect, r)
		}
	}
}

func TestMergeParams(t *testing.T) {
	tests := []struct {
		Base   string
		A      url.Values
		Opts   []MergeOption
		Expect string
		Error  error
	}{
		{
			Base: "https://api.twilio.com/2010-04-01/",
			A: url.Values{
				"a": []string{"b"},
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=b",
		},
		{
			Base:   "https://api.twilio.com/2010-04-01/?a=b",
			A:      url.Values{},
			Expect: "https://api.twilio.com/2010-04-01/?a=b",
		},
		{
			Base:   "https://api.twilio.com/2010-04-01/?a=b",
			A:      nil,
			Expect: "https://api.twilio.com/2010-04-01/?a=b",
		},
		{
			Base: "https://api.twilio.com/2010-04-01/?",
			A: url.Values{
				"a": []string{"b"},
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=b",
		},
		{
			Base: "https://api.twilio.com/2010-04-01/?a=b",
			A: url.Values{
				"a": []string{"b"},
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=b",
		},
		{
			Base: "https://api.twilio.com/2010-04-01/?a=b",
			A: url.Values{
				"a": []string{"b"},
			},
			Opts: []MergeOption{
				Append(true),
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=b&a=b",
		},
		{
			Base: "https://api.twilio.com/2010-04-01/?a=b",
			A: url.Values{
				"a": []string{"b", "c", "d"},
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=b&a=c&a=d",
		},
		{
			Base: "https://api.twilio.com/2010-04-01/?a=b",
			A: url.Values{
				"a": []string{"b", "c", "d"},
			},
			Opts: []MergeOption{
				Append(true),
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=b&a=b&a=c&a=d",
		},
		{
			Base: "https://api.twilio.com/2010-04-01/?a=b&a=c&a=d",
			A: url.Values{
				"a": []string{"c", "d"},
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=c&a=d",
		},
		{
			Base: "https://api.twilio.com/2010-04-01/?a=b&a=c&a=d",
			A: url.Values{
				"a": []string{"c", "d"},
			},
			Opts: []MergeOption{
				Append(true),
			},
			Expect: "https://api.twilio.com/2010-04-01/?a=b&a=c&a=d&a=c&a=d",
		},
	}
	for i, e := range tests {
		r, err := MergeParams(e.Base, e.A, e.Opts...)
		if e.Error != nil {
			fmt.Println("***", err)
			assert.Equal(t, e.Error, err)
		} else if assert.Nil(t, err, fmt.Sprint(err)) {
			fmt.Println("-->", r)
			assert.Equal(t, e.Expect, r, "#%d %v + %v", i, e.Base, e.A)
		}
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
