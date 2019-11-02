package main
import (
    "fmt"
    _ "os"
    _ "bufio"
    _ "io"
    "io/ioutil"
)
func main() {
    // 大文件
    // file, err := os.Open("../test.txt")
    // if err != nil {
    //     fmt.Println("读取失败!")
    //     return
    // }
    // 默认4096
    // reader := bufio.NewReader(file)
    // for {
    //     str, err := reader.ReadString('\n')
    //     if err == io.EOF {
    //         break
    //     }
    //     fmt.Print(str)
    // }
    // defer file.Close()
    // 小文件
    file := "../test.txt"
    content, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("read file err", err)
    }
    fmt.Printf("%v", string(content))
}
