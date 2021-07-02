package main

import (
	"fmt"
	"sync"

	"github.com/h0n9/cert-inspector/cert"
	"github.com/h0n9/cert-inspector/file"
)

func main() {
	// read file
	hostFile := file.NewHostFile("./hosts.yaml")
	err := hostFile.Read()
	if err != nil {
		panic(err)
	}

	// update hostFile.Hosts
	wg := sync.WaitGroup{}
	for i := range hostFile.Hosts {
		h := &hostFile.Hosts[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			cs, err := cert.GetConnState(h)
			if err != nil {
				//fmt.Println(h.Hostname, err)
				return
			}
			if len(cs.PeerCertificates) == 0 {
				fmt.Println(h.Hostname, "couldn't find PeerCertificates")
				return
			}
			pc := cs.PeerCertificates[0]
			// set host data
			h.SetIssuer(pc.Issuer.String())
			h.SetExpiry(pc.NotAfter)
			fmt.Println(h)
		}()
	}

	wg.Wait()

	// save file
	err = hostFile.Save()
	if err != nil {
		panic(err)
	}
}
