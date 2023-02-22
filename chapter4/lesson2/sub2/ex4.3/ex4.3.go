/*
Exercise 4.3
Перепишите функцию reverse так, чтобы вместо среза она использовала указатель на массив.
*/

package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(rotate(a[:], 1))
}

// rotate делает сдвиг элементов на указанное количество позиций
func rotate(slice []int, i int) []int {
	i %= len(slice)                    // получаем крайний элемент заданного сдвига
	tmp := append(slice, slice[:i]...) // добавляем в переменную tmp диапазон значений до крайнего элемента
	copy(slice, tmp[i:])               // копированием отбрасываем значения исходного среза.
	return slice

}
