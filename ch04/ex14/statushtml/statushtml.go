package statushtml

import (
	"html/template"
	"log"
	"os"
	"../github"
)

func CreateIssueHTML() {
	var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>IssueUser</th>
		<th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.IssueUser.HTMLURL}}'>{{.IssueUser.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func CreateUserHTML() {
	var userList = template.Must(template.New("userlist").Parse(`
	<h1>{{.TotalCount}} users</h1>
	<table>
	<tr style='text-align: left'>
		<th>Id</th>
		<th>#</th>
		<th>Avatar</th>
		<th>Score</th>
	</tr>
	{{range .Items}}
	<tr>
		<td>{{.Id}}</td>
		<td><a href='{{.URL}}'>{{.Login}}</a></td>
		<td><img src="{{.AvatarURL}}" height="42" width="42"/></td>
		<td><p>{{.Score}}</p></td>
	</tr>
	{{end}}
	</table>
	`))

	result, err := github.SearchUsers(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := userList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}