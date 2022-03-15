package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("listening...")
    http.HandleFunc("/", helloServer)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

func helloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `<!doctype html>
<html>
<head><title>localhost:8080</title><head>
<body>
Hello!
</body>
</html>`)
}
