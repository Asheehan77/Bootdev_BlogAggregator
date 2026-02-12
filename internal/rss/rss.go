package rss

import (
	"context"
	"net/http"
	"io"
	"encoding/xml"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){

	req,err := http.NewRequestWithContext(ctx,"GET",feedURL,nil)
	if err != nil {
		return nil,err
	}

	req.Header.Set("User-Agent","gator")
	var client http.Client
	
	res,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer res.Body.Close()

	data,err := io.ReadAll(res.Body)
	if err != nil {
		return nil,err
	}

	var rssf RSSFeed
	err = xml.Unmarshal(data,&rssf)
	if err != nil {
		return nil,err
	}


	return &rssf, nil

}
