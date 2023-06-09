# Running Go on Wasm

## Prerequisites

* You have [tinygo](https://tinygo.org/getting-started/install/) (or Go 1.21 when it's available) installed.
* You have a Wasm runtime installed, for example
  [Wasmtime](https://wasmtime.dev/) or
  [WasmEdge](https://wasmedge.org/book/en/quick_start/install.html).

## Configure Go for Wasm

Currently, Go does not support Wasm but TinyGo does. Make sure you have it:

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

## Create, build, and run a HelloWorld app in Wasm

Create `main.go` for a HeloWorld app:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

Build:

```sh
go build main.go
```

Run:

```sh
./main

Hello world!
```

Compile a Wasm binary using TinyGo:

```sh
tinygo build -target=wasi main.go
```

Run with a Wasm runtime such as `wasmtime`:

```sh
wasmtime main.wasm

Hello world!
```

Or `wasmedge`:

```sh
wasmedge main.wasm

Hello world!
```
