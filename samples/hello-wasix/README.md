# Running a Rust HTTP server with Wasix on Wasm

This is a potential Rust HTTP server with [Wasix](https://wasix.org/) on Wasm.
It doesn't work because Wasix is very experimental right now.

## Prerequisites

Install Rust:

```sh
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs |
```

Install Wasmer runtime (Wasix only runs on Wasmer):

```sh
curl https://get.wasmer.io -sSfL | sh -s "v4.0.0-beta.3"
```

## Configure Rust for Wasix

Install Cargo-Wasix:

```sh
cargo install cargo-wasix
```

## Build and run an app with Wasix on Wasm

See [main.rs](./src/main.rs) for an HTTP server in Rust.

Build with Wasix:

```sh
cargo wasix build
```

TODO: This currently fails. Try again some other time.

Run in Wasmer runtime:

```sh
TODO
```

## References

* [Wasix with Axum](https://wasix.org/docs/language-guide/rust/tutorials/wasix-axum)
