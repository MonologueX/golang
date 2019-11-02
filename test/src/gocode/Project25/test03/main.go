package main
import (
    "fmt"
    "bufio"
    "os"
    "io"
)

func main() {
    path := "test.txt"
    // 覆盖
    // file, err := os.OpenFile(path, os.O_WRONLY | os.O_TRUNC, 0666)
    file, err := os.OpenFile(path, os.O_RDWR | os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("open file err", err)
        return
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        str, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        fmt.Print(str)
    }
    str := "你好!\r\n"
    writer := bufio.NewWriter(file)
    for i := 0; i < 5; i++ {
        writer.WriteString(str)
    }
    writer.Flush()
}
