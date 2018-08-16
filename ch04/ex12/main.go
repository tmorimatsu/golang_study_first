package main

import (
	"fmt"
	"log"
	"os"

	"./comics"
)

func main() {
	if len(os.Args) == 1 {
		comics.GetComics("./out.txt")
	} else if len(os.Args) == 2 {
		query := os.Args[1]
		result := comics.Search(query, "./out.txt")
		if len(result) < 1 {
			fmt.Println("no result")
			os.Exit(0)
		}
		for _, num := range result {
			//TODO: api じゃなくてfileから取得
			comic, err := comics.GetComic(num)
			if err != nil {
				log.Fatal(err)
			}
			comics.PrintComic(comic)
		}
		fmt.Println(result)
	} else {
		fmt.Println("follow the right usage")
	}
}
