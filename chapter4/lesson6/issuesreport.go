package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"GolangBook/chapter4/lesson5/github"
)

const templ = `{{.TotalCount}} тем:
{{range .Items}}---------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
