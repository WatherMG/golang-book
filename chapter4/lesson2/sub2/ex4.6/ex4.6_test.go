package main

import (
	"reflect"
	"testing"
)

var testData = []byte("This is  test  text  This is  test  text  This is  test" +
	"  text  This is  test  text              This is  test       text         ")

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert(testData)
	}
}

func Test_convert(t *testing.T) {
	tests := []struct {
		s, want []byte
	}{
		{[]byte("This   is   the   test   text"), []byte("This is the test text")},
	}
	for _, tt := range tests {
		if got := convert(tt.s); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("convert() = %v, want %v", got, tt.want)
		}
	}
}
