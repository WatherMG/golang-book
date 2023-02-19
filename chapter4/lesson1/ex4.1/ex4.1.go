/*
Exercise 4.1
Напишите функцию, которая подсчитывает количество битов, различных в двух дайджестах SHA256
(см. popCount в разделе 2.6.2).
*/

package main

import (
	"crypto/sha256"
	"fmt"

	popcount "GolangBook/chapter2/lesson6/sub2"
)

func countDifferenceBits(sha1, sha2 [sha256.Size]byte) int {
	count := 0
	for i := 0; i < len(sha1); i++ {
		diff := sha1[i] ^ sha2[i]
		count += popcount.PopCount(uint64(diff))
	}
	return count
}

func main() {
	sha1 := sha256.Sum256([]byte("QWERTY"))
	sha2 := sha256.Sum256([]byte("QWERTy"))

	bitsDiff := countDifferenceBits(sha1, sha2)

	fmt.Printf("Количество различающихся битов: %d\nsha1: %x\nsha2: %x\nsha1 == sha2: %t\n",
		bitsDiff, sha1, sha2, sha1 == sha2)
}
