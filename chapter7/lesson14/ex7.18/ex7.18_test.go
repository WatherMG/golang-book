package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	xml := `<doc><h1 class="text">TEXT</h1><p id="p" class="text" style="font: 14px">lorem ipsum</p><div><ul><li id="li active">LI</li><li id="li">LI</li><li id="li">LI</li></ul></div></doc>`
	node, err := parse(strings.NewReader(xml))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(node.(*Element))
	el := node.(*Element)
	expected := `<doc>
  <h1 class="text">
    "TEXT"
  <p id="p" class="text" style="font: 14px">
    "lorem ipsum"
  <div>
    <ul>
      <li id="li active">
        "LI"
      <li id="li">
        "LI"
      <li id="li">
        "LI"`
	got := el.String()[:len(el.String())-2]
	if got != expected[:len(expected)-1] {
		t.Errorf("%q != %q", got, expected[:len(expected)-1])
	}
}
