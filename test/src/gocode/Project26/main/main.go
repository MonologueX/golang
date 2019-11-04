package main
import (
    "fmt"
    "os"
    "flag"
)

// go run main.go -u root -pwd 123456 -h host -port 3306
func main() {
    fmt.Printf("命令行有%v个参数!\n", len(os.Args))
    for i, v := range os.Args {
        fmt.Printf("args[%v]=%v\n", i, v)
    }
    var user string
    var pwd string
    var host string
    var port int
    flag.StringVar(&user, "u", "", "用户名默认为空")
    flag.StringVar(&pwd, "pwd", "", "密码默认为空")
    flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
    flag.IntVar(&port, "port", 3306, "端口号,默认为3306")
    flag.Parse()
    fmt.Printf("user=%q, pwd=%q, host=%q, port=%v\n", user, pwd, host, port)
}
