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
dotnet new wasiconsole -o HelloWasiConsole
```

Build:

```sh
cd HelloWasiConsole
dotnet build
```

Run:

```sh
dotnet run

Running: wasmtime run --dir . -- dotnet.wasm HelloWasiConsole
Using working directory: /Users/atamel/dev/github/meteatamel/wasm-basics/samples/dotnet-wasm/HelloWasiConsole/bin/Debug/net8.0/wasi-wasm/AppBundle
Hello, Wasi Console!
```

You can also run directly with a Wasm runtime such as `wasmtime`:

```sh
cd /Users/atamel/dev/github/meteatamel/wasm-basics/samples/dotnet-wasm/HelloWasiConsole/bin/Debug/net8.0/wasi-wasm/AppBundle
wasmtime run --dir . -- dotnet.wasm HelloWasiConsole

Hello, Wasi Console!
```

<!-- ## Create a single wasm file for the app

So far, we relied on `dotnet.wasm`, a standard build of the .NET runtime for
Wasm to load your and run your apps. Instead, you can create a single Wasm file
to contain the application.

Change the `HelloWasiConsole.csproj` to add `WasmSingleFileBundler`:

```xml
<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
    <RuntimeIdentifier>wasi-wasm</RuntimeIdentifier>
    <OutputType>Exe</OutputType>
    <PublishTrimmed>true</PublishTrimmed>
    <WasmSingleFileBundle>true</WasmSingleFileBundle> 
  </PropertyGroup>
</Project>
-->

## References

* [Experiments with the new WASI workload in .NET 8 Preview
  4](https://youtu.be/gKX-cdqnb8I)