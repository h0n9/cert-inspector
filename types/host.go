package types

type Host struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	Expiry   string `json:"expiry,omitempty"`
}

func NewHost(hostname string, port int) *Host {
	return &Host{
		Hostname: hostname,
		Port:     port,
	}
}
