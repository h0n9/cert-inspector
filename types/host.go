package types

import (
	"fmt"
	"time"
)

const (
	DefaultSSLPort     = 443
	DefaultTimeout     = 10 * time.Second
	DefaultExpWarnDays = 15 * 24 * time.Hour
)

type Host struct {
	Hostname string `json:"hostname" yaml:"hostname"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Issuer   string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
	ExpStr   string `json:"expiry,omitempty" yaml:"expiry,omitempty"`
	expTime  time.Time
}

func NewHost(hostname string, port int) *Host {
	return &Host{
		Hostname: hostname,
		Port:     port,
	}
}

func (h *Host) SetIssuer(issuer string) {
	h.Issuer = issuer
}

func (h *Host) SetExpiry(expTime time.Time) {
	h.expTime = expTime
	h.ExpStr = expTime.Format(time.RFC3339)
}

func (h *Host) String() string {
	// pre-process
	port := h.Port
	if port == 0 {
		port = DefaultSSLPort
	}
	str := ""
	if time.Now().Add(DefaultExpWarnDays).After(h.expTime) {
		str = fmt.Sprintf("\033[1;33m%s\033[0m", "[WARN]\t")
	}
	str += fmt.Sprintf("Hostname: %s\n\tPort: %d\n\tIssuer: %s\n\tExpiry: %s",
		h.Hostname, port, h.Issuer, h.ExpStr)
	return str
}
