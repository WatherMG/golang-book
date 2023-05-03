package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // Не палиндром
		{"desserts", false},   // Полупалиндром
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome возвращает палиндром, длина и содержимое
// которого задаются генератором псевдослучайных чисел rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // Случайная длина до 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // Случайная руна до `\u0999`
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func rndPalindromeWithPunctSpace(rng *rand.Rand) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ" +
		" ,.!? 、。！？")
	n := rng.Intn(25)
	word := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := runes[rng.Intn(len(runes))]
		if unicode.IsLetter(r) {
			word[i] = r
			word[n-1-i] = r
		}
	}
	return string(word)
}

func TestRandomPalindrome(t *testing.T) {
	// Инициализация генератора псевдослучайных чисел.
	seed := time.Now().UTC().UnixNano()
	t.Logf("ГПСЧ инициализирован: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("RandomPalindrome(%q) = false", p)
		}

		pwp := rndPalindromeWithPunctSpace(rng)
		if !IsPalindrome(pwp) {
			t.Errorf("WithPunct(%q) = false", p)
		}
	}
}
