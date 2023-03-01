package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Params struct {
	Owner string
	Repo  string
	Issue
}

// Issue without descriptors doesn't work update issue
type Issue struct {
	Number      int       `json:"number,omitempty"`
	HTMLURL     string    `json:"htmlurl,omitempty"`
	Title       string    `json:"title,omitempty"`
	Body        string    `json:"body,omitempty"`
	State       string    `json:"state,omitempty"`
	StateReason string    `json:"state_reason,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

const baseURL = "https://api.github.com/repos"

var githubToken = os.Getenv("GITHUB_TOKEN")

func getURL(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response error: %s on page: %s\n", resp.Status, resp.Request.URL)
	}
	return resp, nil
}

func putIssue(method string, url string, b io.Reader) bool {
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return false
	}
	req.Header.Set("Authorization", "token "+githubToken)
	req.Header.Set("Content-Type", "application/vnd.github+json")

	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return true
	} else {
		fmt.Fprintf(os.Stderr, "Error with issue #%s.\t", resp.Status)
		return false
	}

}

func (p *Params) GetIssues() (*[]Issue, error) {
	u := fmt.Sprintf("%s/%s/%s/issues", baseURL, p.Owner, p.Repo)

	resp, err := getURL(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}

	return &issues, nil
}

func (p *Params) GetIssue() (*Issue, error) {
	u := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, p.Owner, p.Repo, p.Issue.Number)
	resp, err := getURL(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (p *Params) CreateIssue() bool {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(p.Issue); err != nil {
		return false
	}
	u := fmt.Sprintf("%s/%s/%s/issues", baseURL, p.Owner, p.Repo)

	return putIssue(http.MethodPost, u, buf)
}

func (p *Params) UpdateIssue() bool {
	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(p.Issue); err != nil {
		return false
	}
	u := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, p.Owner, p.Repo, p.Issue.Number)

	return putIssue(http.MethodPatch, u, buf)
}
