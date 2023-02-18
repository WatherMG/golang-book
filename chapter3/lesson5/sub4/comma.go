/*
Example 3.5.3
comma вставляет запятые в строковое представление неотрицательного десятичного числа.
*/

package main

func main() {

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]

}
