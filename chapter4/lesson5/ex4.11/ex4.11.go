/*
Exercise 4.11
Создайте инструмент, который позволит пользователю создавать, читать, обновлять и закрывать темы GitHub из
командной строки, вызывая предпочитаемый пользователем текстовый редактор,
когда требуется ввести текст значительного объема.

Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
invoking their preferred text editor when substantial text input is required.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"GolangBook/chapter4/lesson5/ex4.11/issue"
)

var (
	create = flag.Bool("c", false, "use for creating issue")
	read   = flag.Bool("r", false, "use for reading issue from repo")
	update = flag.Bool("u", false, "use for editing issue")
	list   = flag.Bool("l", false, "use to get list of issues from repo")
	open   = flag.Bool("o", false, "use to open\\reopen issue")
	close_ = flag.Bool("cl", false, "use to close issue")

	owner  = flag.String("owner", "", "")
	repo   = flag.String("repo", "", "")
	number = flag.Int("number", 0, "")
	state  = flag.String("state", "", "open, close")

	title = flag.String("title", "", "title for issue")
	body  = flag.String("body", "", "body for issue")

	githubUser = os.Getenv("GITHUB_USER")
)

func main() {
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if githubUser == "" {
		githubUser = *owner
	}

	flag.Parse()

	switch {
	case *list:
		searchIssues()
	case *read:
		readIssue()
	case *create:
		createIssue()
	case *update:
		updateIssue()
	case *open:
		openIssue()
	case *close_:
		closeIssue()
	}

}
func openEditor(p *issue.Params) *issue.Params {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "notepad"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.CreateTemp("", "issue_temp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	iss, err := p.GetIssue()
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(f).Encode(map[string]string{
		"title": iss.Title,
		"body":  iss.Body,
		"state": iss.State,
	}); err != nil {
		log.Fatal(err)
	}

	// w is parameter for vscode: command to wait for the window to be closed before the command completes
	w := ""
	if editor == "code" {
		w = "-w"
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, f.Name(), w},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	f.Seek(0, 0)
	if err := json.NewDecoder(f).Decode(&p); err != nil {
		log.Fatal(err)
	}
	return p
}

func updateIssue() {

	p := issue.Params{
		Owner: githubUser, Repo: *repo,
		Issue: issue.Issue{
			Title: *title, Body: *body, State: *state, Number: *number,
		},
	}

	openEditor(&p)

	if p.UpdateIssue() {
		fmt.Fprintf(os.Stdout, "\nIssue successfuly updated\n")
	} else {
		fmt.Fprintf(os.Stderr, "\nError editing issue #%d", p.Number)
	}

}

func createIssue() {
	p := issue.Params{
		Owner: githubUser, Repo: *repo,
		Issue: issue.Issue{
			Title: *title, Body: *body, CreatedAt: time.Now(),
		},
	}

	if p.CreateIssue() {
		fmt.Fprintf(os.Stdout, "Issue successfuly created\n")
	} else {
		fmt.Fprint(os.Stderr, "\nError creating 'issue' ")
	}
}

func readIssue() {
	p := issue.Params{Owner: githubUser, Repo: *repo, Issue: issue.Issue{Number: *number}}
	i, err := p.GetIssue()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Printf("#%-5d %s %.60s %s %.20s\n", i.Number, i.Title, i.Body, i.State, i.CreatedAt)
}

func searchIssues() {
	p := issue.Params{Owner: githubUser, Repo: *repo}
	issues, err := p.GetIssues()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	for _, i := range *issues {
		fmt.Printf("#%-5d %s %.60s %s %.20s\n", i.Number, i.Title, i.Body, i.State, i.CreatedAt)
	}
}

func closeIssue() {
	p := issue.Params{Owner: githubUser, Repo: *repo, Issue: issue.Issue{State: "close", Number: *number}}
	if p.UpdateIssue() {
		fmt.Fprintf(os.Stdout, "\nIssue successfuly closed\n")
	} else {
		fmt.Fprintf(os.Stderr, "\nError closing issue #%d", p.Number)
	}
}

func openIssue() {
	p := issue.Params{Owner: githubUser, Repo: *repo, Issue: issue.Issue{State: "open", Number: *number}}
	if p.UpdateIssue() {
		fmt.Fprintf(os.Stdout, "\nIssue successfuly opened\n")
	} else {
		fmt.Fprintf(os.Stderr, "\nError opening issue #%d", p.Number)
	}
}
