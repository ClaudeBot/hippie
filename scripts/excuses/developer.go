package excuses

import (
	"github.com/PuerkitoBio/goquery"
)

const (
	dExcusesAPI = "http://www.devexcuses.com/"
)

type DeveloperExcuse struct{}

func (p *DeveloperExcuse) Run(m []string) (string, error) {
	doc, err := goquery.NewDocument(dExcusesAPI)
	if err != nil {
		return "", err
	}

	excuse := doc.Find(".excuse").First().Text()
	return excuse, nil
}
