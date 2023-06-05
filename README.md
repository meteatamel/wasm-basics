# WebAssembly Basics

![WebAssembly Logo](https://avatars.githubusercontent.com/u/11578470?s=200&v=4)

This repository contains information, references, and samples about running
WebAssembly (Wasm) outside the browser.

## WebAssembly: Why? What?

### Wasm in the browser

[WebAssembly](https://webassembly.org/) (Wasm) is a technology that allows you
to compile code written in over 40+ languages and run it inside sandboxed
environments in a fast, efficient, and safe way. The original use cases were
focused on running native code in web browsers and supported by all major
browsers. For example, there are frameworks like Blazor that enable running C# /
ASP.NET code in the browser running on .NET runtime implemented in WebAssembly.

### Wasm outside the browser

More recently, there’s an attempt to get WebAssembly run outside the browser
with the [WASI](https://wasi.dev/) project. The idea is instead of running apps
as VMs or containers, you can run them as faster, smaller, more secure, and more
portable WebAssembly modules:

* **Faster**: Wasm apps start 10x to 500x faster than containers with no cold-start. 
* **Smaller**: A HelloWorld Rust app in Wasm is 10x smaller than in an OCI container.
* **More secure**: Containers execute in an allow-by-default model whereas Wasm apps execute in a deny-by-default sandbox.
* **More portable**: A container built for `linux/amd64` won’t work on
  `windows/amd64` or even `linux/arm64`. Wasm creates a single `wasm32/wasi`
  module that runs everywhere in a Wasm runtime on any host.

## WASI

<img src="https://wasi.dev/polyfill/WASI-small.png" alt="WASI Logo" width="30%" height="30%">

Wasm code outside of a browser needs a way to talk to the system — a system
interface and that’s what [WebAssembly System Interface](https://wasi.dev/)
(WASI) provides. Wasm runtimes implement WASI.

![WASI architecture](https://github.com/bytecodealliance/wasmtime/raw/main/docs/wasi-software-architecture.png)

A lot of WASI is still in proposals and limited. For example, networking is not
yet part of the WASI standard (so, no socket support in your Wasm module).
However, some Wasm runtimes like `wasmedge` (see [WasmEdge WASI
Socket](https://github.com/second-state/wasmedge_wasi_socket)) and `wasmtime`
implement their own POSIX sockets. There are projects like
[WAGI](https://github.com/deislabs/wagi) to wrap HTTP handlers around Wasi (and
frameworks like [Spin](https://spin.fermyon.dev/) use WAGI) or
[WASIX](https://wasix.org/) to add networking, multi-threading and more to WASI
(but only supported on `wasmer` runtime for now).

## Wasm Runtimes

In Wasm, instead of compiling to an OS and architecture such as `linux/arm64`,
you compile to WebAssembly (`wasm32/wasi`). Wasm has its own bytecode format
that needs to run in a Wasm runtime.

Many wasm runtimes exist (see
[awesome-wasm-runtimes](https://github.com/appcypher/awesome-wasm-runtimes)),
some of the popular ones are:

* [wasmtime](https://github.com/bytecodealliance/wasmtime): A Bytecode Alliance project, designed to run on servers and the cloud.
* [wasmer](https://wasmer.io/): Another popular Wasm runtime from a startup.
* [wasmedge](https://wasmedge.org/): A CNCF project, with more focus on edge
  devices.

## Relevant specs and projects

There are relevant specs and projects (that I haven’t looked in too much
detail):

* [WebAssembly Component Model](https://github.com/WebAssembly/component-model):
  WASI is layered on top of the Component Model, with the Component Model
  providing the foundational building blocks used to define WASI’s interfaces.
  In comparison to traditional Operating Systems, the Component Model fills the
  role of an OS's process model (defining how processes start up and communicate
  with each other) while WASI fills the role of an OS's many I/O interfaces.
* [WAGI](https://github.com/deislabs/wagi): WebAssembly Gateway Interface allows
  you to run WebAssembly WASI binaries as HTTP handlers. Write a command line
  application that prints a few headers, and compile it to `wasm32/wasi`. Add an
  entry to the `modules.toml` matching URL to the Wasm module and that’s it.
* [WASIX](https://wasix.org/): WASIX is the long term stabilization and support
  of the existing WASI plus additional non-invasive syscall extensions (e.g.
  networking, multi-threading) that complete the missing gaps sufficiently
  enough to enable real, practical and useful applications to be compiled and
  used now. It aims to speed up the ecosystem around the WASI so that the
  Wasm’ification of code bases around the world can really start. Only supported
  in `wasmer` runtime right now.

## Running Wasm in containers

There’s ongoing work to run Wasm inside of containers. This might seem
counterintuitive — why take something smaller, faster, and more portable than a
container and run it inside a container? Running Wasm apps inside of a container
gets you the security benefits of the Wasm sandbox and the benefits of the
existing toolchain like Docker, Kubernetes.

### runwasi

<img src="https://raw.githubusercontent.com/containerd/runwasi/main/art/logo/runwasi_icon1.svg" alt="Runwasi Logo" width="30%" height="30%">

Kubernetes relies on a container runtime called
[containerd](https://containerd.io/) (which in turn relies on `runc`) to manage
the lifecycle of containers. [runwasi](https://github.com/deislabs/runwasi) is a
project to integrate Wasm runtimes with `containerd` to enable `containerd` to
manage the lifecycle of Wasm apps. Currently, `wasmtime` and `wasmedge` runtimes
are supported in `runwasi`.
