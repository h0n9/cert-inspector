package cert

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/h0n9/cert-inspector/types"
)

func GetConnState(host *types.Host) (*tls.ConnectionState, error) {
	hostname := host.Hostname
	port := host.Port
	if port == 0 {
		port = types.DefaultSSLPort
	}
	addr := fmt.Sprintf("%s:%d", hostname, port)

	dialer := net.Dialer{Timeout: types.DefaultTimeout}
	conn, err := tls.DialWithDialer(&dialer, "tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.VerifyHostname(hostname)
	if err != nil {
		return nil, err
	}

	cs := conn.ConnectionState()
	return &cs, nil
}
