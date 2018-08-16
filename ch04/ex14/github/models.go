package github

import (
	"time"
)

const baseURL = "https://api.github.com"
const IssuesURL = baseURL + "/search/issues"
const UsersURL = baseURL + "/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	IssueUser     *IssueUser `json:"user"`
	CreateAt time.Time `json:"created_at"`
	Body     string    // md形式
}

type IssueUser struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type UsersSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Users
}

type Users struct {
	Id int
	Login string
	URL string `json:"url"`
	AvatarURL string `json:"avatar_url`
	Type string
	Score float64
}