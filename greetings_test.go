package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Lucas"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Lucas")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Lucas") = %q, %v, quiere coincidencia para %#q, nill`, msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error `, msg, err)
	}
}
