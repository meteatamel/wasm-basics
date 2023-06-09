# WebAssembly outside the browser

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

## Vendors and Frameworks for Wasm+Wasi

These are some notable vendors and frameworks supporting Wasm+Wasi. 

### Docker

Starting in Docker Desktop 4.15, Docker uses `runwasi` to support Wasm workloads running in the following runtimes:

* `wasmedge`
* `wasmtime`
* `spin` from Fermyon
* `slight` from Deislabs

See [Announcing Docker+Wasm Technical Preview 2](https://www.docker.com/blog/announcing-dockerwasm-technical-preview-2/) for more details.

### Azure

Azure AKS is also previewing Wasm node pools using `runwasi`. Currently, there are only `containerd` shims available for [spin](https://spin.fermyon.dev/) and [slight](https://github.com/deislabs/spiderlightning#spiderlightning-or-slight) applications, which use the `wasmtime` runtime.

### Spin

[Spin](https://spin.fermyon.dev/) by Fermyon is a WebAssembly framework for building and running event-driven microservice applications with WebAssembly (Wasm) components. Spin handles the HTTP request/response using WAGI HTTP Executor. Spin SDKs are available in Rust, Go and .NET. and all Wasi-compatible languages are supported. Additionally, [Fermyon Cloud](https://www.fermyon.com/cloud) service fetches source code from the GitHub repo, builds it into Wasm bytecode, runs it as a serverless microservice, and connects HTTP input and output to it using Spin framework. 

### Slight

[SpiderLightning](https://github.com/deislabs/spiderlightning#spiderlightning-or-slight) is a set of WIT interfaces that abstract distributed application capabilities (such as key-value, messaging, HTTP server/client) and a runtime CLI for running Wasm applications that use these capabilities.

## Languages supported in Wasm+WASI

Here’s a [comprehensive list of
languages](https://www.fermyon.com/wasm-languages/webassembly-language-support)
from Fermyon and [The Complex World of Wasm Language
Support](https://www.fermyon.com/blog/complex-world-of-wasm-language-support)
provides the context.

It's not so clear what languages are really supported on Wasm/Wasi outside the
browser at what level. Most languages, such as .NET, are experimental with
Wasm/Wasi support at best and changing constantly. Other languages, like Java,
have some support for Wasm via projects like TeaVM but geared more towards
browsers.

## Samples

These are some samples in this repo:

* [rust-wasm](./samples/rust-wasm/) - Running Rust on Wasm in Docker.
* [dotnet8-wasm](./samples/dotnet8-wasm/) - Running .NET 8 (preview) on Wasm in Docker.
* [go-wasm](./samples/go-wasm/) - Running Go on Wasm.
* [hello-wagi](./samples/hello-wagi) - Running WASI binaries as HTTP handlers with WAGI.
* [python-spin-wasm](./samples/python-spin-wasm) - Running Python with Spin on Wasm.

## References

WASI:

* [https://wasi.dev/](https://wasi.dev/)
* [Standardizing WASI: A system interface to run WebAssembly outside the web](https://hacks.mozilla.org/2019/03/standardizing-wasi-a-webassembly-system-interface/)

Docker + Wasm:

* [Announcing Docker+Wasm Technical Preview 2](https://www.docker.com/blog/announcing-dockerwasm-technical-preview-2/)
* [Introducing the Docker+Wasm Technical Preview](https://www.docker.com/blog/docker-wasm-technical-preview/)
* [Docker + Wasm (Beta)](https://docs.docker.com/desktop/wasm/) 

Azure + Wasm:

* [Create WebAssembly System Interface (WASI) node pools in Azure Kubernetes Service (AKS) to run your WebAssembly (WASM) workload (preview)](https://learn.microsoft.com/en-us/azure/aks/use-wasi-node-pools)

Nigel Poulton’s blog:

* [WebAssembly: The future of cloud computing](https://nigelpoulton.com/webassembly-the-future-of-cloud-computing/)
* [What is cloud native WebAssembly?](https://nigelpoulton.com/what-is-cloud-native-webassembly/)
* [What is runwasi?](https://nigelpoulton.com/what-is-runwasi/)
* [Getting started with Docker + Wasm](https://nigelpoulton.com/getting-started-with-docker-and-wasm/)

Wasm By Example:

* [WASI Introduction](https://wasmbyexample.dev/examples/wasi-introduction/wasi-introduction.all.en-us.html)
* [WASI Hello World](https://wasmbyexample.dev/examples/wasi-hello-world/wasi-hello-world.rust.en-us.html)

WebAssembly for different languages:

* [WebAssembly support in Top 20 languages by
  Fermyon](https://www.fermyon.com/wasm-languages/webassembly-language-support)
* [Enarx - WebAssembly Introduction](https://enarx.dev/docs/WebAssembly/Introduction)
* [WASI support on .NET 8](https://twitter.com/stevensanderson/status/1658845798212202496?s=46&t=qBzme20QIA50uklBQV_ArA): https://twitter.com/stevensanderson/status/1658845798212202496?s=46&t=qBzme20QIA50uklBQV\_ArA
* [The JVM Meets WASI: Writing Cloud-Friendly Wasm Apps Using Java and Friends -
  Joel Dice](https://youtu.be/MFruf7aqcbE)

Talks at Cloud Native Wasm Days:

* [Cloud Native Wasm Day EU 2023](https://colocatedeventseu2023.sched.com/) in [this playlist](https://www.youtube.com/playlist?list=PLj6h78yzYM2Pdj8vnO0wfFyKcbKNy3e5j).
* [Cloud Native Wasm Day North America 2022](https://cloudnativewasmdayna22.sched.com/) in [this playlist](https://www.youtube.com/playlist?list=PLj6h78yzYM2PzLhPvZIihwPShNuXP01C5)
* [Cloud Native Wasm Day EU 2022 ](https://cloudnativewasmdayeu22.sched.com/)in this [playlist](https://www.youtube.com/playlist?list=PLj6h78yzYM2Ni0u-ONljTkv4uOutyjwq9)
