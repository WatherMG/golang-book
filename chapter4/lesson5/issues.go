/*
Example 4.10
Выводит таблицу - результат поиска в GitHub.
*/

package main

import (
	"fmt"
	"log"
	"os"

	"GolangBook/chapter4/lesson5/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d тем:\n", result.TotalCount)
	for _, item := range result.Items {

		fmt.Printf("#%-5d %9.9s %.60s\n",
			item.Number, item.User.Login, item.Title)
	}
}
