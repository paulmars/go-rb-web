package main

import (
  "fmt"
  "net/http"
  "os"
  "code.google.com/p/go.net/websocket"
  "io"
  "net/http"
  "fmt"
  "log"
)

func echoHandler(ws *websocket.Conn) {

  receivedtext := make([]byte, 100)

  n,err := ws.Read(receivedtext)

  if err != nil {
    fmt.Printf("Received: %d bytes\n",n)
  }

  s := string(receivedtext[:n])
  fmt.Printf("Received: %d bytes: %s\n",n,s)

  if _, err := ws.Write([]byte("reboot!\n")); err != nil {
      log.Fatal(err)
  }

  io.Copy(ws, ws)

  fmt.Printf("Sent: %s\n",s)
}

func main() {
  http.Handle("/echo", websocket.Handler(echoHandler))
    http.HandleFunc("/", hello)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
}
