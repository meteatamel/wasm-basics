use std::{thread, time};

fn main() {
    println!("Hello, Wasm before!");
    let three_seconds = time::Duration::from_secs(3);
    thread::sleep(three_seconds);
    println!("Hello, Wasm after!");
}
