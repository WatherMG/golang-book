package main

import (
	"reflect"
	"testing"
)

var (
	testData = []string{"a", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c"}
)

func BenchmarkUniq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unique(testData)
	}
}

func Test_unique(t *testing.T) {
	tests := []struct {
		name, want []string
	}{
		{[]string{"a", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c", "a", "a", "b", "b", "c", "c"}, []string{"a", "b", "c", "a", "b", "c", "a", "b", "c"}},
	}
	for _, tt := range tests {
		if got := unique(tt.name); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("unique() = %v, want %v", got, tt.want)
		}
	}
}
