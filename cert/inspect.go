package cert

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"

	"github.com/h0n9/cert-inspector/types"
)

func Update(host *types.Host) error {
	ips, err := net.LookupIP(host.Hostname)
	if err != nil {
		return err
	}
	cs, err := getConnState(host)
	if err != nil {
		return err
	}

	if len(cs.PeerCertificates) == 0 {
		return fmt.Errorf(host.Hostname, "couldn't find PeerCertificates")
	}

	// set hostost data
	pc := cs.PeerCertificates[0]
	host.SetIPs(ips)
	host.SetIssuer(pc.Issuer.String())
	host.SetExpiry(pc.NotAfter)

	return nil
}

func getConnState(host *types.Host) (*tls.ConnectionState, error) {
	hostname := host.Hostname
	port := host.Port
	if port == 0 {
		port = types.DefaultSSLPort
	}
	addr := net.JoinHostPort(hostname, strconv.Itoa(port))

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
