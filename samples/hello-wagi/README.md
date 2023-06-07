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

## Create a console app and compile for Wasm/Wasi

This console app can be in any language that can compile to Wasm/Wasi. For this
sample, let's use a .NET console app.

```sh
dotnet new console -n HelloWagi
```

Add `Wasi.Sdk` package:

```sh
cd HelloWagi
dotnet add package Wasi.Sdk --prerelease
```

Build:

```sh
dotnet build

  HelloWagi -> /Users/atamel/dev/github/meteatamel/wasm-basics/samples/hello-wagi/HelloWagi/bin/Debug/net8.0/HelloWagi.wasm
```

Run in a Wasm runtime:

```sh
wasmtime ./bin/Debug/net8.0/HelloWagi.wasm

Hello, World!
```

You can also run directly with a Wasm runtime such as `wasmtime`:

```sh
cd /Users/atamel/dev/github/meteatamel/wasm-basics/samples/dotnet-wasm/HelloWasiConsole/bin/Debug/net8.0/wasi-wasm/AppBundle
wasmtime run --dir . -- dotnet.wasm HelloWasiConsole

Hello, Wasi Console!
```

## Run as a WAGI module

To run this sample with WAGI, create a simple `modules.toml` file that maps
paths to a Wasm module:

```
[[module]]
route = "/"
module = "bin/Debug/net8.0/HelloWagi.wasm"
```

Change `Program.cs` to add the following at the beginning of the console app to
print the content type and an empty line. Also change the message to `Hello,
WAGI`:

```csharp
Console.WriteLine("Content-Type: text/plain");
Console.WriteLine();
Console.WriteLine("Hello, WAGI!");
```

Build:

```sh
dotnet build
```

Run as a WAGI module:

```sh
wagi -c modules.toml

Ready: serving on http://127.0.0.1:3000
```

In a seperate terminal, you can hit the url:

```sh
curl http://localhost:3000

Hello, WAGI!
```
