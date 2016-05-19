//go:generate go-bindata -pkg codes -prefix _codes _codes/

package codes

import (
	"encoding/json"
)

type Code struct {
	Code         string `json:"code"`
	Descriptions struct {
		Ietf struct {
			Body string `json:"body"`
			Link string `json:"link"`
		} `json:"ietf"`
		Wikipedia struct {
			Body string `json:"body"`
			Link string `json:"link"`
		} `json:"wikipedia"`
	} `json:"descriptions"`
	References struct {
		Rails struct {
			Title string `json:"title"`
			Value string `json:"value"`
		} `json:"rails"`
	} `json:"references"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
}

type Codes map[string]Code

func All() Codes {
	b, err := Asset("codes.json")
	if err != nil {
		panic(err)
	}

	var codes Codes
	err = json.Unmarshal(b, &codes)
	if err != nil {
		panic(err)
	}

	return codes
}
