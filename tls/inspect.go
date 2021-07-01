package tls

import (
	"crypto/tls"
	"fmt"

	"github.com/h0n9/cert-inspector/types"
)

func GetConnState(host *types.Host) (*tls.ConnectionState, error) {
	addr := fmt.Sprintf("%s:%d", host.Hostname, host.Port)
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.VerifyHostname(host.Hostname)
	if err != nil {
		return nil, err
	}

	connectionState := conn.ConnectionState()
	return &connectionState, nil
}
