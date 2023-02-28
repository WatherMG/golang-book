/*
Exercise 4.10

Измените программу issues так, чтобы она сообщала о результатах с учетом их давности,
деля на категории, например, поданные менее месяца назад, менее года назад и более года.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"GolangBook/chapter4/lesson5/github"
)

var lessMonth []*github.Issue
var lessYear []*github.Issue
var overYear []*github.Issue

func main() {
	result, err := github.SearchIssues([]string{"repo:golang/go is:open json decoder"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d тем:\n", result.TotalCount)
	for _, item := range result.Items {
		t := math.Abs(item.CreatedAt.Sub(time.Now()).Minutes() / 1440)
		if t <= 30 {
			lessMonth = append(lessMonth, item)
		} else if t <= 360 {
			lessYear = append(lessYear, item)
		} else {
			overYear = append(overYear, item)
		}
	}
	fmt.Printf("\nРазмещено меньше месяца назад:\n")
	for _, item := range lessMonth {
		fmt.Printf("\n#%-5d %9.9s %.60s, %.10s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("\nРазмещено меньше года назад:\n")
	for _, item := range lessYear {
		fmt.Printf("\n#%-5d %9.9s %.60s, %.10s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("\nРазмещено больше года назад:\n")
	for _, item := range overYear {
		fmt.Printf("\n#%-5d %9.9s %.60s, %.10s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
