# gsoc-knative 

There are two different dockerfile available 

### v1. Wasm Running in the Go Container

- Runs the WebAssembly module in a Go container with a web server.
- Exposes functionality using JavaScript glue code.
<br>

- Status: go module Error & Range Error at parseJSON
- Dockerfile: hong0331/wasm-app-with-js
- Command: `docker run -p 8080:8080 hong0331/wasm-app-with-js`


![image](https://github.com/iam-zoey/gsoc/assets/67743970/21c89a0c-d211-4bec-a7a2-bd9340253517)



### v2. Wasm Container 

- Runs the WebAssembly module directly in a WebAssembly runtime container.
- Displays parsed JSON by rendering the logs from the container.
<br>

- Status: ✅ Running 
- Dockerfile: hong0331/my-go-app
- Command: `docker run --runtime=io.containerd.wasmtime.v1 hong0331/my-go-app                 `

![image](https://github.com/iam-zoey/gsoc/assets/67743970/2c4c0798-8951-4bae-8ce0-aba8205369d6)
