package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/h0n9/cert-inspector/cert"
	"github.com/h0n9/cert-inspector/file"
)

func run(cmd *cobra.Command, args []string) {
	// get flags
	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		printError("Parse 'update' flag", err)
		return
	}

	// read file
	hostFile := file.NewHostFile("./hosts.yaml")
	err = hostFile.Read()
	if err != nil {
		printError("Read hostFile", err)
		return
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
				printError(h.Hostname, err)
				return
			}

			// print result
			fmt.Println(h)
		}()
	}

	wg.Wait()

	// save file
	if !update {
		return
	}

	err = hostFile.Save()
	if err != nil {
		printError("Save hostFile", err)
		return
	}
}

func printError(where string, err error) {
	fmt.Printf("[ERROR] %s - %s\n", where, err)
}

func main() {
	// get flags
	rootCmd := &cobra.Command{
		Use: "cert-inspector",
		Run: run,
	}
	rootCmd.Flags().BoolP("update", "u", false, "update host file")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
