use std::{thread, time};

fn main() {
    println!("Hello, Wasm before!");
    let duration = time::Duration::from_secs(10);
    thread::sleep(duration);
    println!("Hello, Wasm after!");
}
