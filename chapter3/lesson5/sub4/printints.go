/*
Example 3.5.4
intsToString подобна fmt.Sprint(values), но добавляет запятые
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, _ = fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()

}
