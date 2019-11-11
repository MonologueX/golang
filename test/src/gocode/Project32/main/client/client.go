package main
import (
    "fmt"
    "net"
    "bufio"
    "os"
    "strings"
)
func main() {
    conn, err := net.Dial("tcp", "0.0.0.0:8888")
    if err != nil {
        fmt.Println("client dial err :", err)
        return
    }
    reader := bufio.NewReader(os.Stdin)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("readstring err =", err)
        }
        line  = strings.Trim(line, " \r\n")
        if line == "exit" {
            fmt.Println("client exit!")
            break
        }
        _, err = conn.Write([]byte(line+"\r\n"))
        if err != nil {
            fmt.Println("conn write err =", err)
        }
    }
    fmt.Println("client conn successful, conn =", conn)
}
