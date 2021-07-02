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
			err := cert.Update(h)
			if err != nil {
				fmt.Printf("%s: %s\n", h.Hostname, err)
				return
			}

			// print result
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
