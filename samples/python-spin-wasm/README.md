# Running Python with Spin on Wasm

## Prerequisites

* You have Spin [installed](https://developer.fermyon.com/spin/install).

## Configure Python for Wasm

Install Spin `py2wasm` plugin:

```sh
spin plugins update
spin plugins install py2wasm --yes
```

## Create, build, and run a HelloWorld app with Spin on Wasm

Create a new Spin app with Python template:

```sh
spin new http-py

Enter a name for your new application: hello_wasm
Description: Python Spin HelloWorld app
HTTP base: /
HTTP path: /...
```

Build:

```sh
cd hello_wasm
spin build
```

Run:

```sh
spin up

Serving http://127.0.0.1:3000
Available Routes:
  hello-wasm: http://127.0.0.1:3000 (wildcard)
```

In a separate terminal, you can hit the url:

```sh
curl http://127.0.0.1:3000

Hello from the Python SDK
```

## References

* [Taking Spin for a spin](https://developer.fermyon.com/spin/quickstart)
* [Building Spin Components in Python](https://developer.fermyon.com/spin/python-components)
