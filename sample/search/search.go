package search

import (
	"fmt"
	"log"
)

type Result struct {
	Code    int
	Message int
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) (result Result, err error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- Result) (err error) {
	var searchResults Result
	searchResults, err = matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	results <- searchResults
	return
}

func Display(results chan Result) {
	for res := range results {
		fmt.Println("content %v", res)
	}
}
