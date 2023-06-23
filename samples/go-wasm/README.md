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

Create [hello-wasm.go](./hello-wasm.go) to access the filesystem:

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

## Running HTTP in Wasm+Wasi

You might think that you can compile and run any Go app to Wasm+Wasi but this is
not true. Wasi preview1 does not support sockets yet. 

To show this, create a simple HelloWorld HTTP server in [hello-http.go](./hello-http.go):

```go
package main

import (
  "fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  log.Println("Request received:", r.Method, r.URL.Path)
  fmt.Fprint(w, "Hello World!")
}

func main() {
  http.HandleFunc("/", handler)
  log.Println("Server started. Listening on :8080...")

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatalf("ListenAndServe error:%s ", err.Error())
  }
}
```

Build with Go:

```sh
GOOS=wasip1 GOARCH=wasm go1.21rc2 build -o hello-http.wasm hello-http.go
```

Run with `wasmtime`:

```sh
wasmtime run hello-http.wasm
```

And you get an error because sockets are not supported yet:

```sh
2023/06/23 10:41:06 Server started. Listening on :8080...
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
net.(*fakeNetFD).accept(...)
        /Users/atamel/sdk/go1.21rc2/src/net/net_fake.go:231
net.(*netFD).accept(0x144e2c0)
        /Users/atamel/sdk/go1.21rc2/src/net/fd_wasip1.go:88 +0x44
net.(*TCPListener).accept(0x142a0e0)
        /Users/atamel/sdk/go1.21rc2/src/net/tcpsock_posix.go:152 +0x4
net.(*TCPListener).Accept(0x142a0e0)
        /Users/atamel/sdk/go1.21rc2/src/net/tcpsock.go:315 +0x8
net/http.(*Server).Serve(0x1474000, {0xc2138, 0x142a0e0})
        /Users/atamel/sdk/go1.21rc2/src/net/http/server.go:3056 +0x30
net/http.(*Server).ListenAndServe(0x1474000)
        /Users/atamel/sdk/go1.21rc2/src/net/http/server.go:2985 +0x10
net/http.ListenAndServe(...)
        /Users/atamel/sdk/go1.21rc2/src/net/http/server.go:3239
main.main()
        /Users/atamel/dev/github/meteatamel/wasm-basics/samples/go-wasm/hello-http.go:18 +0xe
```
