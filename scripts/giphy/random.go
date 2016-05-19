package giphy

import (
	"net/url"
)

type RandomData struct {
	ID               string `json:"id"`
	URL              string `json:"url"`
	ImageOriginalURL string `json:"image_original_url"`
	ImageURL         string `json:"image_url"`
	ImageMP4URL      string `json:"image_mp4_url"`
}

type Random struct {
	Data RandomData `json:"data"`
}

func (r *Random) ImageURL() string {
	return r.Data.ImageURL
}

func (r *Random) Run(m []string) (string, error) {
	query := m[1]
	p := url.Values{}
	p.Set("tag", query)
	_, err := giphyRequest("random", p, r)
	if err != nil {
		return "", err
	}

	return r.ImageURL(), nil
}
