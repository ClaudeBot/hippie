package giphy

import (
	"math/rand"
	"net/url"
	"time"
)

type Search struct {
	Data []DefaultData `json:"data"`
}

func (s *Search) ImageURL() string {
	// Pretty much like the random API ...
	rand.Seed(time.Now().Unix())
	return s.Data[rand.Intn(len(s.Data))].Images.Original.URL
}

func (s *Search) Run(m []string) (string, error) {
	query := m[1]
	p := url.Values{}
	p.Set("q", query)
	_, err := giphyRequest("search", p, s)
	if err != nil {
		return "", err
	}

	return s.ImageURL(), nil
}
