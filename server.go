package main

import (
    "fmt"
    "net/http"
    "os"
    "bufio"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Onica!! Final version\n")

    file,_ := os.Open("hostname/hostname.txt")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
            fmt.Fprintf(w,"Hostname: " + scanner.Text())
    }
}
