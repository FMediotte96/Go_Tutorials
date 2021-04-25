package greetings

import (
	"regexp"
	"testing"
)

//TestHelloName calls greetings.Hello with a name, checking
//for a valid return value
func TestHelloName(t *testing.T) {
	name := "Facundo"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Facundo")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Facundo") = %q, %v, want math for %#q, nil`, msg, err, want)
	}
}

//TestHelloEmpty calls greetings.Hello with an empty string,
//checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
