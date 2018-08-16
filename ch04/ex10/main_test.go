package main

import (
	"testing"
	"time"
	"log"
	"./github"
)

func TestCompareDate1(t *testing.T) {
	t.Log("timeの順序が正常な時、trueを返す")
	formerTime, latterTime := time.Now(), time.Now().Add(time.Hour*10)
	expected := true
	actual := compareDate(formerTime, latterTime)
	if expected != actual {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestCompareDate2(t *testing.T) {
	t.Log("timeの順序が逆の時、falseを返す")
	formerTime, latterTime := time.Now(), time.Now().AddDate(-1,0,0)
	expected := false
	actual := compareDate(formerTime, latterTime)
	if expected != actual {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestSortByCreateAt1(t *testing.T) {
	t.Log("CreateAtが早い順に並び替える")
	time1, err := time.Parse("2006-01-02 15:04:05 MST", "2018-01-15 11:31:20 UTC")
	if err != nil {
		log.Fatal(err)
	}
	time2, err := time.Parse("2006-01-02 15:04:05 MST", "2014-08-11 11:50:20 UTC")
	if err != nil {
		log.Fatal(err)
	}
	time3, err := time.Parse("2006-01-02 15:04:05 MST", "2013-11-15 21:31:20 UTC")
	if err != nil {
		log.Fatal(err)
	}
	git1 := github.Issue{CreateAt:time1}
	git2 := github.Issue{CreateAt:time2}
	git3 := github.Issue{CreateAt:time3}
	gits := []github.Issue{git1, git2, git3}
	expected := []github.Issue{git3, git2, git1}
	actual := sortByCreateAt(gits)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
	}
}

// これは基本必要ないが違うパターンも年のため試す目的で残している
func TestSortByCreateAt2(t *testing.T) {
	t.Log("CreateAtが早い順に並び替える")
	time1, err := time.Parse("2006-01-02 15:04:05 MST", "2012-01-15 11:31:20 UTC")
	if err != nil {
		log.Fatal(err)
	}
	time2, err := time.Parse("2006-01-02 15:04:05 MST", "2014-08-11 11:50:20 UTC")
	if err != nil {
		log.Fatal(err)
	}
	time3, err := time.Parse("2006-01-02 15:04:05 MST", "2013-11-15 21:31:20 UTC")
	if err != nil {
		log.Fatal(err)
	}
	git1 := github.Issue{CreateAt:time1}
	git2 := github.Issue{CreateAt:time2}
	git3 := github.Issue{CreateAt:time3}
	gits := []github.Issue{git1, git2, git3}
	expected := []github.Issue{git1, git3, git2}
	actual := sortByCreateAt(gits)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("actual %v\nwant %v", actual, expected)
		}
	}
}