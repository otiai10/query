package query

import (
	. "github.com/otiai10/mint"
	"net/http"
	"testing"
)

func TestStrings(t *testing.T) {
	r := req("?key_foo=jack+john+jessie")
	foos := Strings(r, "key_foo", []string{"this", "is", "default"})
	Expect(t, foos).ToBe([]string{"jack", "john", "jessie"})
}

func req(querystring string) *http.Request {
	req, _ := http.NewRequest("GET", "http://google.com"+querystring, nil)
	return req
}
