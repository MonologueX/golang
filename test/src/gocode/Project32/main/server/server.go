package main
import (
    "fmt"
    "net"
)

func process(conn net.Conn) {
    defer conn.Close()
    for {
        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println(conn.RemoteAddr().String(), "client exit!")
            return
        }
        fmt.Print(string(buf[:n]))
    }
}

func main() {
    fmt.Println("服务器开始监听!")
    listen, err := net.Listen("tcp", "0.0.0.0:8888")
    if err != nil {
        fmt.Println("listen err!")
        return
    }
    defer listen.Close()
    fmt.Println("监听成功!")
    fmt.Printf("listen suc = %v\n", listen)
    for {
        // 等待客户端连接
        fmt.Println("等待客户端连接!")
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("err=%v, accept err!\n", err)
            continue
        } else {
            fmt.Printf("conn=%v, accept successful!\n", conn)
            fmt.Printf("client ip:%v\n", conn.RemoteAddr().String())
        }
        go process(conn)
    }
}
