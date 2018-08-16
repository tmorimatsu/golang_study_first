package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	items := []github.Issue{}
	for _, item := range result.Items {
		items = append(items, *item)
	}
	sortByCreateAt(items)
	month, year := false, false
	for _, item := range items {
		// FixMe: calendarに沿った1ヶ月を利用するように
		// FixMe: タグを付けるほうが良さそう
		if !year && time.Now().Sub(item.CreateAt).Hours() < 8760.0 {
			year = true
			fmt.Println("ここより上は一年以上古い")
		}
		if !month && time.Now().Sub(item.CreateAt).Hours() < 720.0 {
			month = true
			fmt.Println("ここより上は一ヶ月以上古い")
		}
		fmt.Printf("#%-5d %v %.55s %v\n", item.Number, item.User.Login, item.Title, item.CreateAt)
	}
}


func sortByCreateAt(items []github.Issue) []github.Issue {
	// FixMe: クソソートの改修
	for i, _ := range items {
		for j, _ := range items {
			if i > j {
				continue
			}
			if !compareDate(items[i].CreateAt, items[j].CreateAt) {
				tmp := items[i]
				items[i] = items[j]
				items[j] = tmp
			}
		}
	}
	return items
}

// ２つのDateを正しい順序であるかどうかを返却します
// formerDate が latterDate より早いものであればtrue
func compareDate(formerDate time.Time, latterDate time.Time) bool {
	return latterDate.Sub(formerDate) > 0
}