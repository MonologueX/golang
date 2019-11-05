package main
import (
    "fmt"
)
func main() {
    intChan := make(chan int, 10)
    for i := 0; i < cap(intChan); i++ {
        intChan<- i*2
    }
    // 关闭管道，没有关闭读完还会读取
    close(intChan)
    // 遍历
    for v := range intChan {
        fmt.Println("v =", v)
    }
}
