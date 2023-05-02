/*Package word представляет утилиты для игр со словами.*/
package word

// IsPalindrome сообщает, является ли s палиндромом.
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
