# Running Rust on Wasm

## Prerequisites

* You have Rust [installed](https://www.rust-lang.org/tools/install).
* You have a Wasm runtime installed, for example
  [Wasmtime](https://wasmtime.dev/) or
  [WasmEdge](https://wasmedge.org/book/en/quick_start/install.html).

## Create, build, and run a HelloWorld app

Create:

```sh
cargo new hello-wasm
```

Build:

```sh
cd hello-wasm
cargo build
```

Run:

```sh
./target/debug/hello-wasm

Hello, world!
```

## Configure Rust to compile to Wasm

Add `wasm32-wasi` target:

```sh
rustup target add wasm32-wasi
```

## Build and run HelloWorld app in Wasm

Change the `main.rs` to access the filesystem:

```rust
use std::io::prelude::*;
use std::fs;

fn main() {
    println!("Hello, Wasm!");

    // Create a file
    // We are creating a `helloworld.txt` file in the `/helloworld` directory
    // This code requires the Wasi host to provide a `/helloworld` directory on the guest.
    // If the `/helloworld` directory is not available, the unwrap() will cause this program to panic.
    // For example, in Wasmtime, if you want to map the current directory to `/helloworld`,
    // invoke the runtime with the flag/argument: `--mapdir /helloworld::.`
    // This will map the `/helloworld` directory on the guest, to  the current directory (`.`) on the host
    let mut file = fs::File::create("/helloworld/helloworld.txt").unwrap();

    // Write the text to the file we created
    write!(file, "Hello world!\n").unwrap();
}
```

Build for Wasm/Wasi:

```sh
cargo build --target wasm32-wasi
```

Run in a Wasm runtime such as `wasmtime` giving it access to the current
directory on host to `/helloworld` directory on the guest:

```sh
wasmtime --mapdir /helloworld::. target/wasm32-wasi/debug/hello-wasm.wasm

Hello, Wasm!
```

You should also see a `helloworld.txt` file created.

Or in `wasmedge`:

```sh
wasmedge --dir /helloworld:. target/wasm32-wasi/debug/hello-wasm.wasm

Hello, Wasm!
```

## References

* [Wasm by Example - Rust](https://wasmbyexample.dev/examples/wasi-hello-world/wasi-hello-world.rust.en-us.html)
