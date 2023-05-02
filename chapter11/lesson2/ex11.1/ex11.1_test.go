package charcount

import (
	"strings"
	"testing"
)

var tests = []struct {
	input string
	want  [5]int
}{
	{"a", [5]int{0, 1, 0, 0, 0}},
	{"asd", [5]int{0, 3, 0, 0, 0}},
	{"qwerty", [5]int{0, 6, 0, 0, 0}},
	{"й", [5]int{0, 0, 1, 0, 0}},
	{"цукен", [5]int{0, 0, 5, 0, 0}},
	{"Привет, мир!", [5]int{0, 3, 9, 0, 0}},
	{"Hello, World!", [5]int{0, 13, 0, 0, 0}},
	{"Γεια σου, κόσμε!", [5]int{0, 4, 12, 0, 0}},
	{"你好，世界!", [5]int{0, 1, 0, 5, 0}},
	{"안녕하세요, 세상!", [5]int{0, 3, 0, 7, 0}},
	{"Bonjour à tous!", [5]int{0, 14, 1, 0, 0}},
	{"Merhaba, Dünya!", [5]int{0, 14, 1, 0, 0}},
	{"萨沙走在公路上，吮吸着烘干机", [5]int{0, 0, 0, 14, 0}},
}

func TestGetCharCount(t *testing.T) {
	for _, test := range tests {
		if got := GetCharCount(strings.NewReader(test.input)); got != test.want {
			t.Errorf("GetCharCount(%q) = %v", test.input, got)
		}
	}
}
