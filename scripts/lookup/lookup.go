package ping

import (
	"github.com/claudebot/hippie/lambda"
	"net"
	"strings"
)

type ReverseLookup struct{}
type HostLookup struct{}

// type IPLookup struct{}

func init() {
	lambda.Register("(?i)^/lookup ((?:[0-9]{1,3}.){3}[0-9]{1,3})$", &ReverseLookup{})
	lambda.Register("(?i)^/lookup-host (.+)$", &HostLookup{})
	// lambda.Register("(?i)^/lookup-ip (.+)$", &IPLookup{})
}

func (p *ReverseLookup) Run(m []string) (string, error) {
	addr := m[1]
	names, err := net.LookupAddr(addr)
	// TODO: return error message (but, this should probably be addressed
	// in the actual handler)
	if err != nil {
		return "", err
	}
	return strings.Join(names, ", "), nil
}

func (p *HostLookup) Run(m []string) (string, error) {
	host := m[1]
	addrs, err := net.LookupHost(host)
	if err != nil {
		return "", err
	}
	return strings.Join(addrs, ", "), nil
}

// func (p *IPLookup) Run(m []string) (string, error) {
// 	host := m[1]
// 	ips, err := net.LookupIP(host)
// 	if err != nil {
// 		return "", err
// 	}

// 	var sips []string
// 	for _, ip := range ips {
// 		sips = append(sips, ip.String())
// 	}

// 	return strings.Join(sips, ", "), nil
// }
