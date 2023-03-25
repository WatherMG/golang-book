package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "1234567890"
	var limit int64 = 3
	b := &bytes.Buffer{}
	lr := LimitReader(strings.NewReader(s), limit)
	n, err := b.ReadFrom(lr)
	if n != limit || err != nil {
		t.Errorf("len(s)=%d, limit=%d", n, limit)
	}
	if b.String() != s[:limit] {
		t.Errorf("'%s' != '%s'", b.String(), s[:limit])
	}
}
