# Running .NET 8 (preview) on Wasm

## Prerequisites

* You have .NET 8 Preview 4 [installed](https://dotnet.microsoft.com/en-us/download/dotnet/8.0).
* You have a Wasm runtime installed, for example
  [Wasmtime](https://wasmtime.dev/) or
  [WasmEdge](https://wasmedge.org/book/en/quick_start/install.html).

## Configure .NET for Wasm

Make sure you have .NET 8 Preview 4:

```sh
dotnet --version

8.0.100-preview.4.23260.5
```

Install `wasi-experimental` workload that has a new `Wasi Console App` template:

```sh
dotnet workload install wasi-experimental
```

## Create, build, and run a HelloWorld app in Wasm

Make sure the new `Wasi Console App` template is installed:

```sh
dotnet new list wasi

These templates matched your input: 'wasi'

Template Name     Short Name   Language  Tags
----------------  -----------  --------  ----------------
Wasi Console App  wasiconsole  [C#]      Wasi/WasiConsole
```

Create a new `Wasi Console App` using the template:

```sh
dotnet new wasiconsole -o HelloWasm
```

Build:

```sh
cd HelloWasm
dotnet build
```

Run:

```sh
dotnet run

Running: wasmtime run --dir . -- dotnet.wasm HelloWasm
Using working directory: /Users/atamel/dev/github/meteatamel/wasm-basics/samples/dotnet-wasm/HelloWasm/bin/Debug/net8.0/wasi-wasm/AppBundle
Hello, Wasi Console!
```

You can also run directly with a Wasm runtime such as `wasmtime`:

```sh
cd /Users/atamel/dev/github/meteatamel/wasm-basics/samples/dotnet-wasm/HelloWasm/bin/Debug/net8.0/wasi-wasm/AppBundle
wasmtime run --dir . -- dotnet.wasm HelloWasm

Hello, Wasi Console!
```

## Change app to access filesytem

Chane `Program.cs` to access the filesystem:

```csharp
using System;
using System.IO;

Console.WriteLine("Hello, Wasm!");

// Create a file
// We are creating a `helloworld.txt` file in the `/helloworld` directory
// This code requires the Wasi host to provide a `/helloworld` directory on the guest.
// If the `/helloworld` directory is not available, the `File.WriteAllText()` will fail.
// For example, in Wasmtime, if you want to map the current directory to `/helloworld`,
// invoke the runtime with the flag/argument: `--mapdir /helloworld::.`
// This will map the `/helloworld` directory on the guest, to  the current directory (`.`) on the host
string path = "/helloworld/helloworld.txt";
string content = "Hello world!\n";
using (StreamWriter sw = File.CreateText(path))
{
    sw.Write(content);
}

Console.WriteLine("Created helloworld.txt");
```

## Create a single Wasm file for the app

So far, we relied on `dotnet.wasm`, a standard build of the .NET runtime for
Wasm to load your and run your apps. Instead, you can create a single Wasm file
to contain the application.

Change the `HelloWasm.csproj` to the following:

```xml
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
    <!-- <RuntimeIdentifier>wasi-wasm</RuntimeIdentifier> -->
    <OutputType>Exe</OutputType>
    <!-- <PublishTrimmed>true</PublishTrimmed> -->
  </PropertyGroup>
</Project>
```

Add the `Wasi.Sdk` package:

```sh
dotnet add package Wasi.Sdk --prerelease
```

Build for Wasm+Wasi:

```sh
dotnet build

  HelloWasm -> /Users/atamel/dev/github/meteatamel/wasm-basics/samples/dotnet8-wasm/HelloWasm/bin/Debug/net8.0/HelloWasm.wasm
```

Run in a Wasm runtime such as `wasmtime`:

```sh
wasmtime --mapdir /helloworld::. bin/Debug/net8.0/HelloWasm.wasm

Hello, Wasm!
Created helloworld.txt
```

You can try another Wasm runtime like `wasmedge`:

```sh
wasmedge --dir /helloworld:. bin/Debug/net8.0/HelloWasm.wasm

Hello, Wasm!
Created helloworld.txt
```

## References

* [Experiments with the new WASI workload in .NET 8 Preview
  4](https://youtu.be/gKX-cdqnb8I)
* [Bringing WebAssembly to the .NET Mainstream - Steve Sanderson,
  Microsoft](https://youtu.be/PIeYw7kJUIg?list=PLj6h78yzYM2Ni0u-ONljTkv4uOutyjwq9)
* [Experimental WASI SDK for .NET
  Core](https://github.com/SteveSandersonMS/dotnet-wasi-sdk#how-to-use-aspnet-core-applications)
