package main

import (
    "fmt"
    "errors"
)

func test() {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("err =", err)
        }
    } ()
    num1 := 1
    num2 := 0
    res := num1 / num2
    fmt.Println("res =", res)
}

// 自定义
func readConf(name string) (err error) {
    if name == "config.ini" {
        return nil
    } else {
        return errors.New("读取文件错误")
    }
}

func test2() {
    err := readConf("config.ini")
    if err != nil {
        panic(err)
    }
    fmt.Println("...")
}

func main() {
    test2()
    test()
}
