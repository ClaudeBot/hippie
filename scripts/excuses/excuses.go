package excuses

import (
	"github.com/claudebot/hippie/lambda"
)

func init() {
	lambda.Register("(?i)^/excuse$", &ProgrammingExcuses{})
	lambda.Register("(?i)^/dev-excuse$", &DeveloperExcuses{})
}
