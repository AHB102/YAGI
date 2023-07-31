package core

import (
	"net"
)

// ReverseDNS performs a reverse DNS lookup for the given IP address
func ReverseDNS(ipAddress string) ([]string, error) {
	names, err := net.LookupAddr(ipAddress)
	if err != nil {
		return nil, err
	}
	return names, nil
}
