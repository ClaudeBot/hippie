package excuses

import (
	"github.com/claudebot/hippie/lambda"
)

func init() {
	lambda.Register("(?i)^/excuse$", &ProgrammingExcuse{})
	lambda.Register("(?i)^/dev-excuse$", &DeveloperExcuse{})
}
