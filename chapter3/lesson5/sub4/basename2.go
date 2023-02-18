/*
Example 3.5.2
basename(s) убирает из s все префиксы, имеющие вид пути файловой системы с компонентами, разделенными косыми чертами,
а также суффикс, представляющий тип файла
basename убирает компоненты каталога и суффикс типа файла.
Более простая версия использует библиотечную функцию strings.Lastlndex
а => a, a.go => a, a/b/c.go => с, a/b.c.go => b.c
*/

package main

import (
	"strings"
)

func main() {

}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1, если "/" не найден
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
