package cert

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/h0n9/cert-inspector/types"
)

const (
	DefaultSSLPort = 443
	DefaultTimeout = 3 * time.Second
)

func GetConnState(host *types.Host) (*tls.ConnectionState, error) {
	hostname := host.Hostname
	port := host.Port | DefaultSSLPort
	addr := fmt.Sprintf("%s:%d", hostname, port)

	dialer := net.Dialer{Timeout: DefaultTimeout}
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
