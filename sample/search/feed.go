package search

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

const dataFile = "data/data.json"

func RetrieveFeeds() (feeds []*Feed, err error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return
	}

	err = json.NewDecoder(file).Decode(&feeds)
	return
}
