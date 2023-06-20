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