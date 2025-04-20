package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

const randomURL = "https://c.xkcd.com/random/comic"

func XKCD(event *handler.CommandEvent) error {
	comic, err := getXKCD()
	if err != nil {
		return fmt.Errorf(`getting comic: %w`, err)
	}

	msg := fmt.Sprintf("[%s: %d](%s)", comic.Title, comic.Number, comic.URL)
	return event.CreateMessage(discord.MessageCreate{
		Content: msg,
	})
}

type ComicXKCD struct {
	URL    string `json:"-"`
	Title  string `json:"safe_title"`
	Number int    `json:"num"`
	Image  string `json"img"`
}

func getXKCD() (*ComicXKCD, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(randomURL)
	if err != nil {
		return nil, err
	}

	url, ok := resp.Header["Location"]
	if !ok {
		return nil, errors.New("could not get comic")
	}

	resp, err = http.Get(url[0] + "info.0.json")
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(resp.Body)
	comic := ComicXKCD{
		URL: url[0],
	}
	err = json.Unmarshal(bytes, &comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil

}
