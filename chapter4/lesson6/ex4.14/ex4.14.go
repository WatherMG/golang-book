/*
Exercise 4.14

Создайте веб-сервер, однократно запрашивающий информацию у GitHub, а затем позволяющий выполнять навигацию
по списку сообщений об ошибках, контрольных точек и пользователей.

Create a web server that queries GitHub once and then allows navigation of the list of bug reports, milestones, and users.
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"GolangBook/chapter4/lesson6/ex4.14/issues"
)

var p = "chapter4/lesson6/ex4.14/templates"
var issueListTemplate = template.Must(template.New("issueList.html").ParseFiles(p + "/issueList.html"))
var issueTemplate = template.Must(template.ParseFiles(p + "/issue.html"))
var milestoneTemplate = template.Must(template.ParseFiles(p + "/milestone.html"))
var userTemplate = template.Must(template.ParseFiles(p + "/user.html"))

var cache *issues.Cache
var searchQuery string

func init() {
	cache = &issues.Cache{
		IssuesByID:     make(map[int]*issues.Issue),
		MilestonesByID: make(map[int]*issues.Milestone),
		Users:          make(map[string]*issues.User),
	}
}

// run in browser type localhost:8000/?q=repo:golang/go%20is:open%20json%20decoder
// or localhost:8000/?q=SEARCH_QUERY
// To prevent html templates from generating an error, you need to check the path (var p)
func main() {
	http.HandleFunc("/", getIssuesList)
	http.HandleFunc("/issue/", getIssue)
	http.HandleFunc("/milestone/", getMilestone)
	http.HandleFunc("/user/", getUser)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getIssuesList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	if q != "" {
		if cache.Issues == nil || searchQuery != q { // for reuse cache in same query when refresh
			searchQuery = q
			res, err := issues.SearchIssues(q)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			cache.Issues = res.Items
			for _, issue := range res.Items {
				cache.IssuesByID[issue.Number] = issue
				if issue.User != nil {
					u, err := issue.User.GetUser()
					if err != nil {
						log.Fatal(err)
					}
					cache.Users[issue.User.Login] = u
				}
				if issue.Milestone != nil {
					cache.MilestonesByID[issue.Milestone.Number] = issue.Milestone
				}
			}
		}

		if err := issueListTemplate.Execute(w, cache); err != nil {
			log.Fatal(err)
		}
	}
}

func getIssue(w http.ResponseWriter, r *http.Request) {
	num := getID(w, r)
	issue, ok := cache.IssuesByID[num]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("Issue not found: %d", num)))
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := issueTemplate.Execute(w, issue); err != nil {
		log.Fatal(err)
	}
}

func getMilestone(w http.ResponseWriter, r *http.Request) {
	num := getID(w, r)
	issue, ok := cache.MilestonesByID[num]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("Issue not found: %d", num)))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := milestoneTemplate.Execute(w, issue); err != nil {
		log.Fatal(err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	name := pathParts[2]
	if err := len(name); err == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("User %s not found", name)))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	user, ok := cache.Users[name]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("User %s not found", user)))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := userTemplate.Execute(w, user); err != nil {
		log.Fatal(err)
	}

}

func getID(w http.ResponseWriter, r *http.Request) int {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	numStr := pathParts[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Issue number isn't a number %s", numStr)))
		if err != nil {
			log.Fatal(err)
		}
	}
	return num
}
