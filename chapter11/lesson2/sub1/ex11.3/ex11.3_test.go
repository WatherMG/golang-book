package word

import (
	"math/rand"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func notPalindrome(rng *rand.Rand) string {
	n := rng.Intn(21) + 4 // Случайная длина до 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := letterRunes[rand.Intn(len(letterRunes))]
		runes[i] = r
		runes[n-1-i] = r
	}
	p := string(runes)
	for string(runes) == p {
		r := letterRunes[rand.Intn(len(letterRunes))]
		pos := rng.Intn(n / 2)
		runes[rng.Intn(2)*((n-1)-2*pos)+pos] = r
	}
	return string(runes)
}

func TestRandomWord(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("ГПСЧ инициализирован: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		np := notPalindrome(rng)
		if IsPalindrome(np) {
			t.Errorf("RandomWord(%q) = true", np)
		}
	}
}
