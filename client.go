package main

import (
  "code.google.com/p/go.net/websocket"
  // "io"
  "log"
  // "net/http"
  "fmt"
)

func main() {
    origin := "https://murmuring-springs-9063.herokuapp.com/"
    url := "ws://murmuring-springs-9063.herokuapp.com:80/echo"

    ws, err := websocket.Dial(url, "", origin)
    if err != nil {
        log.Fatal(err)
    }

    if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
        log.Fatal(err)
    }

    var msg = make([]byte, 512)
    var n int
    if n, err = ws.Read(msg); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Received: %s.\n", msg[:n])
}
