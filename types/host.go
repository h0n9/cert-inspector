package types

import (
	"fmt"
)

type Host struct {
	Hostname string `json:"hostname" yaml:"hostname"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Issuer   string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
	Expiry   string `json:"expiry,omitempty" yaml:"expiry,omitempty"`
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

func (h *Host) SetExpiry(expiry string) {
	h.Expiry = expiry
}

func (h *Host) String() string {
	return fmt.Sprintf("Hostname: %s, Port: %d, Issuer: %s, Expiry: %s",
		h.Hostname, h.Port, h.Issuer, h.Expiry)
}
