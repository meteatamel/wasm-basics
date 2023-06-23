# Running Go on Wasm

## Prerequisites

* You have [Go 1.21 RC](https://go.dev/blog/go1.21rc) or [tinygo](https://tinygo.org/getting-started/install/) installed.
* You have a Wasm runtime installed, for example
  [Wasmtime](https://wasmtime.dev/) or
  [WasmEdge](https://wasmedge.org/book/en/quick_start/install.html).

## Configure Go for Wasm

Currently, Go does not support Wasm but Go 1.21 RC and TinyGo does. I show how
to use both.

Make sure you have Go 1.21 RC installed:

```sh
go1.21rc2 version

go version go1.21rc2 darwin/amd64
```

Or if you prefer TinyGo, make sure you have it:

```sh
tinygo version

tinygo version 0.27.0 darwin/amd64 (using go version go1.20.5 and LLVM version 15.0.0)
```

When you first compile with TinyGo, you might get this error:

```sh
error: could not find wasm-opt, set the WASMOPT environment variable to override
```

This means, you need to install and/or link `binaryen` (at least on MacOS)

```sh
brew install binaryen
brew link binaryen
```

## Create, build, and run an app in Wasm+Wasi

Create `hello-wasm.go` to access the filesystem:

```go
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
```

Build with Go:

```sh
GOOS=wasip1 GOARCH=wasm go1.21rc2 build -o hello-wasm.wasm hello-wasm.go
```

Or, build with TinyGo:

```sh
tinygo build -target=wasi hello-wasm.go
```

Run in a Wasm runtime such as `wasmtime`:

```sh
wasmtime --mapdir /helloworld::. hello-wasm.wasm

Hello, Wasm!
Created helloworld.txt
```

Or, run in another Wasm runtime such as `wasmedge`:

```sh
wasmedge --dir /helloworld:. hello-wasm.wasm

Hello, Wasm!
Created helloworld.txt
```
