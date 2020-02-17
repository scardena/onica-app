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
    fmt.Fprintf(w, "Hello, Onica!!\n")

    file,_ := os.Open("hostname.txt")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
            fmt.Fprintf(w,"Hostname: " + scanner.Text())
    }
}