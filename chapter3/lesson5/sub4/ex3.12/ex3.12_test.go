package anagram

import (
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		f, s string
		want bool
	}{
		{"Statue of Liberty", "Built to stay free", true},
		{"Statue Liberty", "Built to stay free", false},
		{"eat", "tea", true},
		{"listen", "silent", true},
		{"anagram", "nag a ram", true},
		{"Elvis", "lives", true},
		{"A gentleman", "elegant man", true},
		{"Clint Eastwood", "old west action", true},
		{"Tom Marvolo Riddle", "I am Lord Voldemort", true},
		{"dormitory", "dirty room", true},
		{"the eyes", "they see", true},
		{"slot machines", "cash lost in me", true},
		{"debit card", "bad credit", true},
		{"astronomer", "moon starer", true},
		{"tea", "coffee", false},
		{"Statue Liberty", "Built to stay free", false},
		{"hello", "world", false},
		{"мама мыла раму", "раму мыла мама", true},
		{"воз и ныне там", "там и ныне воз", true},
		{"я с миром", "мир со мной", false},
		{"была цель, жить правильно", "жить было правильно, цель аль", false},
		{"нам дали тепло", "тепло дали нам", true},
		{"отвертка лежит рядом с телом", "отвертка рядом лежит с телом", true},
		{"возможности", "положительный", false},
		{"честный человек", "мудрость", false},
		{"чувство ответственности", "ответственность чувство", false},
		{"тяжело в учении - легко в бою", "в учении легко - тяжело в бою", true},
		{"надо отдать должное", "отдать надо должное", true},
	}
	for _, tt := range tests {
		got := isAnagram(tt.f, tt.s)
		if got != tt.want {
			t.Errorf("isAnagram(%q, %q), got %v, want(%v)", tt.f, tt.s, got, tt.want)
		}
	}

}
