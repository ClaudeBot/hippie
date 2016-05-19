package excuses

import (
	"github.com/PuerkitoBio/goquery"
)

const (
	pExcusesAPI = "http://programmingexcuses.com/"
)

type ProgrammingExcuses struct{}

func (p *ProgrammingExcuses) Run(m []string) (string, error) {
	doc, err := goquery.NewDocument(pExcusesAPI)
	if err != nil {
		return "", err
	}

	return doc.Find(".wrapper a").First().Text(), nil
}
