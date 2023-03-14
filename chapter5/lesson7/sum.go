/*
Example 5.11
*/

package main

import "fmt"

func main() {
	fmt.Println(sum())           // 0
	fmt.Println(sum(3))          // 3
	fmt.Println(sum(1, 2, 3, 4)) // 10
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // 10

	fmt.Printf("%T\n", f) // func(...int)
	fmt.Printf("%T\n", g) // func([]int)

}
func f(...int) {}
func g([]int)  {}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
