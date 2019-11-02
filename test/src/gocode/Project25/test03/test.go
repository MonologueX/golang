package main
import (
    "fmt"
    "io/ioutil"
)

func main() {
    path1 := "test.txt"
    path2 := "testCopy.txt"
    data, err := ioutil.ReadFile(path1)
    if err != nil {
        fmt.Println("read file err", err)
        return
    }
    err = ioutil.WriteFile(path2, data, 0666)
    if err != nil {
        fmt.Println("write file err", err)
    }
}
