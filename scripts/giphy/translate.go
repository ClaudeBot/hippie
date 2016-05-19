package giphy

import (
	"net/url"
)

type Translate struct {
	Data DefaultData `json:"data"`
}

func (t *Translate) ImageURL() string {
	return t.Data.Images.Original.URL
}

func (t *Translate) Run(m []string) (string, error) {
	query := m[1]
	p := url.Values{}
	p.Set("s", query)
	_, err := giphyRequest("translate", p, t)
	if err != nil {
		return "", err
	}

	return t.ImageURL(), nil
}
