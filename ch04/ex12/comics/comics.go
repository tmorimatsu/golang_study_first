package comics

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string
}

func GetComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.StatusCode)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}

func GetComics(path string) {
	// 総数を先に取得したい
	for i := 1; i < 300; i++ {
		// coChannel := make(chan Comic error)
		// go func() { coChannel <- GetComic(i) }()
		comic, err := GetComic(i)
		if err != nil {
			continue
		}
		writeComics(comic, path)
	}
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := 0
	start := 0
	stop := 0
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if unicode.IsLetter(r) {
			start = i - size
			break
		}
	}
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if unicode.IsLetter(r) {
			stop = i - size
			break
		}
	}
	if stop > start {
		token = data[start:stop]
	}
	return i, token, nil
}

func Search(query, path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 4096)
	var numbers []int
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		comic := new(Comic)
		json.Unmarshal(line, comic)
		if strings.Count(string(line), query) > 0 {
			numbers = append(numbers, comic.Num)
		}
	}
	return numbers
}

func writeComics(comic Comic, path string) {
	by, err := json.Marshal(&comic)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// ioutil.WriteFile(path, by, os.FileMode(0600))
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(string(by) + "\n")
}

func PrintComic(comic Comic) {
	// by, err := json.Marshal(&comic)
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }
	fmt.Println(comic.Num)
	fmt.Println()
	fmt.Println(comic.Year)
	fmt.Println(comic.Month)
	fmt.Println(comic.Day)
	fmt.Println()
	fmt.Println(comic.Title)
	fmt.Println()
	fmt.Println(comic.Transcript)
	fmt.Println()
	fmt.Println(comic.Alt)
	fmt.Println()
	fmt.Println(comic.Img)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
