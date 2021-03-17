package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run() {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// Create an unbuffered channel to receive match results to display.
	results := make(chan Result)
	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			err = Match(matcher, feed, feed.Site, results)
			if err != nil {
				log.Fatal(err)
			}
			waitGroup.Done()
		}(matcher, feed)
	}

	//等待完成
	go func() {
		waitGroup.Wait()
		close(results)
	}()
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
