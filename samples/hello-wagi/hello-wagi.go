package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Content-Type: text/plain")
	fmt.Println("Status: 200")
	fmt.Println()
	fmt.Println("Hello WAGI from Go!")

	// Headers are placed in environment variables
	envVars := os.Environ()
	fmt.Printf("### Environment variables: %d ###\n", len(envVars))
	for _, envVar := range envVars {
		fmt.Println(envVar)
	}

	// Query parameters are sent in as command line options
	args := os.Args[1:]
	fmt.Printf("### Query parameters: %d ###\n", len(args))
	for _, arg := range args {
		fmt.Printf("Argument=%s\n", arg)
	}

	// Incoming HTTP payloads are sent in via STDIN
	fmt.Println("### HTTP payload ###")
	payload, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading payload:", err)
		return
	}
	fmt.Println(string(payload))
}
