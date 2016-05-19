package excuses

import (
	"github.com/PuerkitoBio/goquery"
)

const (
	dExcusesAPI = "http://www.devexcuses.com/"
)

type DeveloperExcuses struct{}

func (p *DeveloperExcuses) Run(m []string) (string, error) {
	doc, err := goquery.NewDocument(dExcusesAPI)
	if err != nil {
		return "", err
	}

	return doc.Find(".excuse").First().Text(), nil
}
