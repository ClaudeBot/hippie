package excuses

import (
	"github.com/PuerkitoBio/goquery"
)

const (
	pExcusesAPI = "http://programmingexcuses.com/"
)

type ProgrammingExcuse struct{}

func (p *ProgrammingExcuse) Run(m []string) (string, error) {
	doc, err := goquery.NewDocument(pExcusesAPI)
	if err != nil {
		return "", err
	}

	excuse := doc.Find(".wrapper a").First().Text()
	return excuse, nil
}
