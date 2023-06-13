// Based on
// https://wasmbyexample.dev/examples/wasi-hello-world/wasi-hello-world.rust.en-us.html

// Import rust's io and filesystem module
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
