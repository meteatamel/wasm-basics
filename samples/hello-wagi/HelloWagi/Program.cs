Console.WriteLine("Content-Type: text/plain");
Console.WriteLine("Status: 200");
Console.WriteLine();
Console.WriteLine("Hello WAGI from C#!");

// Headers are placed in environment variables
var envVars = Environment.GetEnvironmentVariables();
Console.WriteLine($"### Environment variables: {envVars.Keys.Count} ###");
foreach (var variable in envVars.Keys)
{
    Console.WriteLine($"{variable} = {envVars[variable]}");
}

// Query parameters, when present, are sent in as command line options
// TODO: This doesn't work due to Wasi.Sdk package
// Console.WriteLine($"### Query parameters: {args.Length} ###");
// if (args.Length > 0)
// {
//     foreach (var arg in args)
//     {
//         Console.WriteLine($"Argument={arg}");
//     }
// }

// Incoming HTTP payloads are sent in via STDIN
// TODO: This doesn't work due to Wasi.Sdk package
// Console.WriteLine($"### HTTP payload ###");
// string payload = Console.ReadLine();
// Console.WriteLine(payload);