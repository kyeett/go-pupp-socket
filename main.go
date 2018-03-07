package main

import (
    "flag"
    "log"
    "net/http"
    // "github.com/gorilla/websocket"
)

var (
    addr = flag.String("addr", ":8081", "http service addr")
)

func serveHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }
}

func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        log.Fatal("filename not specified")
    }
    //    filename := flag.Args()[0]
    http.HandleFunc("/", serveHome)
    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal(err)
    }
}
