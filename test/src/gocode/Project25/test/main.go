package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    path := "test.txt"
    file, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("open file err", err)
        return
    }
    defer file.Close()
    str := "hello golang!\n"
    writer := bufio.NewWriter(file)
    for i := 0; i < 10; i++ {
        writer.WriteString(str)
    }
    writer.Flush()
}
