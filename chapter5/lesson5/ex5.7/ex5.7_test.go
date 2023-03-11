package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestName(t *testing.T) {
	input := `
<html><head><title>Test</title><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"></head><body><img src="img.png"><img src="img2.jpg"><p>This is a test paragraph.</p></body></html>
`
	output := `<html>
  <head>
    <title>
      Test
    </title>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta http-equiv="X-UA-Compatible" content="ie=edge"/>
  </head>
  <body>
    <img src="img.png"/>
    <img src="img2.jpg"/>
    <p>
      This is a test paragraph.
    </p>
  </body>
</html>
`

	buf := &bytes.Buffer{}

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Errorf("Cant parse HTML: %v", err)
	}
	forEachNode(doc, startElement, endElement, buf)

	if buf.String() != output {
		t.Errorf("prettyPrint() want = %s, got = %s", output, buf.String())
	}
}
