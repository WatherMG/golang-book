/*
Exercise 4.14

Создайте веб-сервер, однократно запрашивающий информацию у GitHub, а затем позволяющий выполнять навигацию
по списку сообщений об ошибках, контрольных точек и пользователей.

Create a web server that queries GitHub once and then allows navigation of the list of bug reports, milestones, and users.
*/

package issues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const URL = "https://api.github.com/search/issues"
const userURL = "https://api.github.com/users/"

var githubToken = os.Getenv("GITHUB_TOKEN")

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Body      string
	User      *User
	Milestone *Milestone
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	PublicRepos int    `json:"public_repos,omitempty"`
	Followers   int    `json:"followers,omitempty"`
	Name        string `json:"name,omitempty"`
	Login       string `json:"login,omitempty"`
	Type        string `json:"type,omitempty"`
	Location    string `json:"location,omitempty"`
	Email       string `json:"email,omitempty"`
	BIO         string `json:"bio,omitempty"`
	HTMLURL     string `json:"html_url"`
}

type Milestone struct {
	Number      int
	Title       string
	Description string
}

type Cache struct {
	Issues         []*Issue
	Users          map[string]*User
	IssuesByID     map[int]*Issue
	MilestonesByID map[int]*Milestone
}

// SearchIssues запрашивает issues у Github
func SearchIssues(terms string) (*IssueSearchResult, error) {
	q := url.QueryEscape(terms)
	resp, err := http.Get(URL + "?q=" + q)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Необходимо закрыть resp.Body на всех путях выполнения
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error: %s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil

}

func (u *User) GetUser() (*User, error) {
	q := url.QueryEscape(u.Login)
	uri := fmt.Sprintf("%s%s", userURL, q)

	client := &http.Client{}

	buf := &bytes.Buffer{}

	req, err := http.NewRequest(http.MethodGet, uri, buf)
	req.Header.Set("Authorization", "token "+githubToken)
	req.Header.Set("Content-Type", "application/vnd.github+json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("response error: %s user %s", resp.Status, u.Login)
	}
	var result User
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &result, nil
}

func (p *Issue) IssueURL() string {
	return fmt.Sprintf("/issue/%d", p.Number)
}
func (p *Issue) MilestoneURL() string {
	return fmt.Sprintf("/milestone/%d", p.Milestone.Number)
}
func (u *User) UserURL() string {
	return fmt.Sprintf("/user/%s", u.Login)
}
func (p *Issue) FormatDate() string {
	f := p.CreatedAt.Format("02.01.2006 | ")
	return f + strconv.Itoa(int(time.Since(p.CreatedAt).Hours()/24)) + " ago"
}
