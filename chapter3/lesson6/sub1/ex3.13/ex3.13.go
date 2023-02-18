/*
Exercise 3.13
Напишите объявления const для КВ, MB и так далее до YB настолько компактно, насколько сможете.
*/

package main

const (
	B  = 1
	KB = B * 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)
