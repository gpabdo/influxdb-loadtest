package main
import (
    "fmt"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("udp", "127.0.0.1:8092")

    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }

    // Number of messages per second.
    var messages int = 3000

    ticker := time.NewTicker(time.Nanosecond * time.Duration(1000000000 / messages) )

    for range ticker.C {
      fmt.Fprintf(conn, "responseTime,verb=GET,url=/,statusCode=200 timeDelta=5,clientIP=\"192.168.1.1\"") 
    }

    conn.Close()
}
