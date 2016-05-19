package httpcodes

import (
	"fmt"
	"github.com/claudebot/hippie/lambda"
	"github.com/claudebot/hippie/scripts/httpcodes/codes"
)

var (
	statusCodes = codes.All()
)

type HTTPCodes struct{}

func init() {
	lambda.Register("(?i)^/http (.+)$", &HTTPCodes{})
}

func (h *HTTPCodes) Run(m []string) (string, error) {
	code := m[1]
	v, ok := statusCodes[code]
	if !ok {
		return "", fmt.Errorf("HTTP status code could not be found: %s", code)
	}

	return fmt.Sprintf("%s (%s)", v.Descriptions.Ietf.Body, v.Descriptions.Ietf.Link), nil
}
