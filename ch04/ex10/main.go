package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/github"
)

WIP

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9 %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
