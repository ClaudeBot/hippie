package giphy

import (
	"encoding/json"
	"github.com/claudebot/hippie/lambda"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	giphyAPI = "http://api.giphy.com/v1/gifs/"
)

type Giphyer interface {
	ImageURL() string
}

type DefaultData struct {
	ID     string `json:"id,omitempty"`
	Slug   string `json:"slug,omitempty"`
	Rating string `json:"rating,omitempty"`
	Images struct {
		Original struct {
			URL string `json:"url"`
		} `json:"original"`
	} `json:"images"`
}

type Search struct {
	Data []DefaultData `json:"data"`
}

type Translate struct {
	Data DefaultData `json:"data"`
}

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

func giphyRequest(e string, p url.Values, v interface{}) (*http.Response, error) {
	base, err := url.Parse(giphyAPI)
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(e)
	if err != nil {
		return nil, err
	}

	api := base.ResolveReference(endpoint)

	if p == nil {
		p = url.Values{}
	}
	p.Set("api_key", "dc6zaTOxFJmzC")
	api.RawQuery = p.Encode()

	// TODO: handle 404, 500, etc
	res, err := http.Get(api.String())
	if err != nil {
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return nil, err
	}

	// NOTE: res.Body is drained, don't reuse
	return res, err
}

func init() {
	lambda.Register("(?i)^/giphy (.+)$", &Search{})
	lambda.Register("(?i)^/translate (.+)$", &Translate{})
	lambda.Register("(?i)^/random (.+)$", &Random{})
}

func (s *Search) ImageURL() string {
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
