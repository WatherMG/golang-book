/*
Exercise 5.19
Воспользуйтесь функциями panic и recover для написания функции,
которая не содержит инструкцию return, но возвращает ненулевое значение.

Use panic and recover to write a function that contains no return statement yet returns a non-zero value.
*/

package main

import "fmt"

func main() {
	fmt.Println(returnRecover())
}

// returnRecover возвращает значение после восстановления из паники
// без использования `return`
func returnRecover() (result int, err error) {
	defer func() {
		if p := recover(); p != nil {
			result = 42
			err = fmt.Errorf("recovered msg: %v", p)
		}
	}()
	panic("It's panic!")
}
