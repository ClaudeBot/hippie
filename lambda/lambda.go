package lambda

import (
	"regexp"
	"sync"
)

var (
	registryMu sync.RWMutex
	registry   = make(map[*regexp.Regexp]Lambda)
)

type Lambda interface {
	Run([]string) (string, error)
}

func Register(p string, l Lambda) {
	registryMu.Lock()
	defer registryMu.Unlock()
	if len(p) == 0 {
		panic("lambda: Register regexp is nil")
	}
	if l == nil {
		panic("lambda: Register lambda is nil")
	}
	re := regexp.MustCompile(p)
	registry[re] = l
}

func Run(msg string) (string, error) {
	for re, l := range registry {
		if matches := re.FindStringSubmatch(msg); matches != nil {
			return l.Run(matches)
		}
	}
	return "", nil
}
