package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, Wasm!")

	// Create a file
	// We are creating a `helloworld.txt` file in the `/helloworld` directory
	// This code requires the Wasi host to provide a `/helloworld` directory on the guest.
	// If the `/helloworld` directory is not available, the `ioutil.WriteFile()` will fail.
	// For example, in Wasmtime, if you want to map the current directory to `/helloworld`,
	// invoke the runtime with the flag/argument: `--mapdir /helloworld::.`
	// This will map the `/helloworld` directory on the guest, to  the current directory (`.`) on the host
	err := ioutil.WriteFile("/helloworld/helloworld.txt", []byte("Hello world!\n"), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created helloworld.txt")
}
