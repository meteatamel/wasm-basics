# Running WASI binaries as HTTP handlers with WAGI

WebAssembly Gateway Interface ([WAGI](https://github.com/deislabs/wagi)) allows
you to run WebAssembly WASI binaries as HTTP handlers. Write a command line
application that prints a few headers, and compile it to `wasm32-wasi`. Add an
entry to the `modules.toml` matching URL to Wasm module. That's it. You can use
any programming language that can compile to `wasm32-wasi`.

Headers are placed in environment variables. Query parameters, when present, are
sent in as command line options. Incoming HTTP payloads are sent in via STDIN.
And the response is simply written to STDOUT.

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

Change [Program.cs](./HelloWagi/Program.cs) to print the following:

```csharp
Console.WriteLine("Content-Type: text/plain");
Console.WriteLine("Status: 200");
Console.WriteLine();
Console.WriteLine("Hello WAGI from C#!");

// Headers are placed in environment variables
var envVars = Environment.GetEnvironmentVariables();
Console.WriteLine($"### Environment variables: {envVars.Keys.Count} ###");
foreach (var variable in envVars.Keys)
{
    Console.WriteLine($"{variable} = {envVars[variable]}");
}
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
Status: 200

Hello WAGI from C#!
### Environment variables: 1 ###
DOTNET_SYSTEM_GLOBALIZATION_INVARIANT = true
```

## Create a Go app for WAGI

For the next sample, let's use a Go console app.

Create [hello-wagi.go](./hello-wagi.go) to print the following:

```go
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
### Environment variables: 23 ###
SERVER_PORT = 3000
REQUEST_METHOD = GET
X_MATCHED_ROUTE = /csharp
...
```

```sh
curl -X POST http://127.0.0.1:3000/csharp\?arg1=value1 -d 'Hello World'

Hello WAGI from Go!
### Environment variables: 24 ###
HTTP_CONTENT_LENGTH=11
X_RAW_PATH_INFO=
REMOTE_USER=
PATH_INFO=
PATH_TRANSLATED=
HTTP_USER_AGENT=curl/7.88.1
SCRIPT_NAME=/go
...
### Query parameters: 1 ###
Argument=arg1=value1
### HTTP payload ###
Hello World
```

## References

* [Wasm, WASI, Wagi: What are they?](hhttps://www.fermyon.com/blog/wasm-wasi-wagi)
