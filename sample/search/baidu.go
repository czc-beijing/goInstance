package search

import (
	"encoding/json"
	"log"
)

import "github.com/idoubi/goz"

type BaiDu struct {
}

func init() {
	var matcher BaiDu
	Register("baidu", matcher)
}

type Resp struct {
	Status      string      `json:"status"`
	Msg         string      `json:"msg"`
	City        string      `json:"city"`
	Weathercode string      `json:"weathercode"`
	Update      string      `json:"update"`
	Data        interface{} `json:"data"`
}

//goland:noinspection ALL
func (b BaiDu) Search(feed *Feed, searchTerm string) (result Result, err error) {
	cli := goz.NewClient()
	resp, err := cli.Get(feed.Link)
	if err != nil {
		log.Fatalln(err)
	}
	var baiduResp Resp
	body, err := resp.GetBody()
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &baiduResp)
	if err != nil {
		return
	}
	result.Code = 1
	return
}
