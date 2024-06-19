# gsoc-knative 

There are two different dockerfile available 

### v1. hong0331/wasm-app-with-js
- Status: go module Error & Range Error at parseJSON
  <br> 
```docker run -p 8080:8080 hong0331/wasm-app-with-js ```

### v2. hong0331/my-go-app 
It displays parsed JSON by rendering the logs from the container. 
In the current `main.go` under func-go>cnd>fhttp, uncomment the line 45 ~ 81 to run knative with this image

- Status: working
  <br> 
```docker run --runtime=io.containerd.wasmtime.v1 hong0331.my-go-app```

