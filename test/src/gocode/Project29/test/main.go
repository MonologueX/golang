package main
import (
    "fmt"
)

func writeData(intChan chan int) {
    for i := 0; i < 100; i++ {
        intChan<- i
        fmt.Println("write data", i)
    }
    close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
    for {
        v, ok := <-intChan
        if !ok {
            break
        }
        fmt.Printf("读到数据%v\n", v)
    }
    exitChan<- true
    close(exitChan)
}
func main() {
    intChan := make(chan int, 100)
    exitChan := make(chan bool, 1)
    go writeData(intChan)
    go readData(intChan, exitChan)
    for {
        _, ok := <-exitChan
        if !ok {
            break
        }
    }
}
