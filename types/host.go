package types

import (
	"fmt"
	"time"

	"github.com/h0n9/cert-inspector/util"
)

const (
	DefaultSSLPort     = 443
	DefaultTimeout     = 10 * time.Second
	DefaultExpWarnDays = 15
)

type Host struct {
	Hostname string `json:"hostname" yaml:"hostname"`
	IPAddr   string `json:"ip_addr" yaml:"ip_addr"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Issuer   string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
	ExpDate  string `json:"exp_date,omitempty" yaml:"exp_date,omitempty"`
	ExpDays  int    `json:"exp_days,omitempty" yaml:"exp_days,omitempty"`
	expTime  time.Time
}

func NewHost(hostname string, port int) *Host {
	return &Host{
		Hostname: hostname,
		Port:     port,
	}
}

func (h *Host) SetIPAddr(ipAddr string) {
	h.IPAddr = ipAddr
}

func (h *Host) SetIssuer(issuer string) {
	h.Issuer = issuer
}

func (h *Host) SetExpiry(expTime time.Time) {
	h.expTime = expTime
	h.ExpDate = expTime.Format(time.RFC3339)
	h.ExpDays = int(time.Since(expTime).Hours()/24) * -1
}

func (h *Host) String() string {
	// pre-process
	port := h.Port
	if port == 0 {
		port = DefaultSSLPort
	}
	str := ""

	if h.ExpDays > DefaultExpWarnDays {
		str += util.Info("[GOOD]\t")
	} else {
		str += util.Warn("[WARN]\t")
	}

	str += fmt.Sprintf("Hostname: %s\n"+
		"\tIP Address: %s\n"+
		"\tPort: %d\n"+
		"\tIssuer: %s\n"+
		"\tExpiry Date: %s\n"+
		"\tDays left for expiry: %d",
		h.Hostname, h.IPAddr, port, h.Issuer, h.ExpDate, h.ExpDays)
	return str
}
