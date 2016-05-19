package giphy

import (
	"encoding/json"
	"github.com/claudebot/hippie/lambda"
	"net/http"
	"net/url"
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
