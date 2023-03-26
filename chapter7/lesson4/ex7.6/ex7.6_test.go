package tempconv

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"testing"
)

var temp = CelsiusFlag("temp", 20.0, "temperature")

func TestCelsiusFlag(t *testing.T) {
	var tcs = []struct {
		args   []string
		expect string
	}{
		{[]string{"testflag", "-temp", "0K"}, "-273.15°C\n"},
		{[]string{"testflag", "-temp", "0°K"}, "-273.15°C\n"},
		{[]string{"testflag", "-temp", "0C"}, "0.00°C\n"},
		{[]string{"testflag", "-temp", "0°C"}, "0.00°C\n"},
		{[]string{"testflag", "-temp", "32F"}, "0.00°C\n"},
		{[]string{"testflag", "-temp", "32°F"}, "0.00°C\n"},
	}
	for _, tc := range tcs {
		os.Args = tc.args
		stdout = new(bytes.Buffer)
		flag.Parse()
		fmt.Fprintln(stdout, *temp)
		actual := stdout.(*bytes.Buffer).String()
		if actual != tc.expect {
			t.Errorf("Args: %v, Expects: %v, Actual: %v", tc.args, tc.expect, actual)
		}
	}
}
