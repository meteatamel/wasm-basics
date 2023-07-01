FROM scratch
COPY ./target/wasm32-wasi/debug/hello-wasm.wasm /hello-wasm.wasm
ENTRYPOINT [ "/hello-wasm.wasm" ]
