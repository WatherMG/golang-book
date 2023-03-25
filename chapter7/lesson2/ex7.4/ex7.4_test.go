package main

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	s := "Hello, World"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Errorf("n=%d, err=%s", n, err)
	}
	if b.String() != s {
		t.Errorf("'%s' != '%s'", b.String(), s)
	}
}

func TestURLReader(t *testing.T) {
	s := "<html><head><meta/></head><body><div><p>Hello, World!</p></div></body></html>"
	_, err := html.Parse(NewReader(s))
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
