package main

import (
	"html/template"
	"log"
	"os"

	"GolangBook/chapter4/lesson5/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	report, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(os.Stdout, report); err != nil {
		log.Fatal(err)
	}
}
