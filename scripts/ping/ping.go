package ping

import (
	"github.com/claudebot/hippie/lambda"
)

type Ping struct{}

func init() {
	lambda.Register("(?i)^/ping$", &Ping{})
}

func (p *Ping) Run(m []string) (string, error) {
	return "PONG", nil
}
