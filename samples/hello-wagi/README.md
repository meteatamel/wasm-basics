# Running WASI binaries as HTTP handlers with WAGI

WebAssembly Gateway Interface ([WAGI](https://github.com/deislabs/wagi)) allows
you to run WebAssembly WASI binaries as HTTP handlers. Write a command line
application that prints a few headers, and compile it to `wasm32-wasi`. Add an
entry to the `modules.toml` matching URL to Wasm module. That's it. You can use
any programming language that can compile to `wasm32-wasi`.

## Prerequisites

* You have a Wasm runtime installed, for example
  [Wasmtime](https://wasmtime.dev/) or
  [WasmEdge](https://wasmedge.org/book/en/quick_start/install.html).

## Install WAGI

Get the latest from here: https://github.com/deislabs/wagi/releases

Untar:

```sh
tar -zxf wagi-v0.8.1-macos-amd64.tar.gz
```

Move to `/usr/local/bin`:

```sh
sudo mv ./wagi /usr/local/bin/wagi
```

## Create a .NET app for WAGI

For this sample, let's use a .NET console app.

```sh
dotnet new console -n HelloWagi
```

Add `Wasi.Sdk` package:

```sh
cd HelloWagi
dotnet add package Wasi.Sdk --prerelease
```

Change `Program.cs` to print the content type and an empty line (required for
WAGI) and also change the message:

```csharp
Console.WriteLine("Content-Type: text/plain");
Console.WriteLine();
Console.WriteLine("Hello WAGI from C#!");
```

Build:

```sh
dotnet build

  HelloWagi -> /Users/atamel/dev/github/meteatamel/wasm-basics/samples/hello-wagi/HelloWagi/bin/Debug/net8.0/HelloWagi.wasm
```

Run in a Wasm runtime:

```sh
wasmtime ./bin/Debug/net8.0/HelloWagi.wasm

Content-Type: text/plain

Hello WAGI from C#!
```

## Create a Go app for WAGI

For this sample, let's use a Go console app.

Create `hello-wagi.go` to print the content type and an empty line (required for
WAGI) and also a message:

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Content-Type: text/plain");
  fmt.Println();
  fmt.Println("Hello WAGI from Go!")
}
```

Build:

```sh
tinygo build -target=wasi hello-wagi.go
```

Run in a Wasm runtime:

```sh
wasmtime ./hello-wagi.wasm

Content-Type: text/plain

Hello WAGI from Go!
```

## Run as a WAGI module

To run these samples with WAGI, create a simple `modules.toml` file that maps
paths to Wasm modules:

```yaml
[[module]]
route = "/csharp"
module = "HelloWagi/bin/Debug/net8.0/HelloWagi.wasm"

[[module]]
route = "/go"
module = "hello-wagi.wasm"
```

Run as a WAGI module:

```sh
wagi -c modules.toml

Ready: serving on http://127.0.0.1:3000
```

In a separate terminal, you can curl different paths to hit different modules:

```sh
curl http://localhost:3000/csharp

Hello WAGI from C#!

curl http://localhost:3000/go

Hello WAGI from Go!
```
