package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", handler1)
    http.ListenAndServe("192.168.31.168:8008", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, getHtmlFile("index1.html"))
}

func getHtmlFile(path string) (fileHtml string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    rd := bufio.NewReader(file)
    for {
        line, err := rd.ReadString('\n')
        if err != nil || io.EOF == err {
            break
        }
        fileHtml += line
    }
    return fileHtml
}
