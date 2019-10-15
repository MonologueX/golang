package main

import (
    "fmt"
)

var (
    Fun1 = func(n1 int, n2 int) int {
        return n1 * n2
    }
)

func main() {
    // 匿名函数
    // 使用一次
    res1 := func (n1 int, n2 int) int {
        return n1 + n2
    } (10, 20)
    fmt.Println("res1=", res1)

    // 使用多次
    sub := func (n1 int, n2 int) int {
        return n1 - n2
    }
    res2 := sub(20, 10)
    fmt.Println("res2=", res2)

    // 全局
    res3 := Fun1(4, 9)
    fmt.Println("res3=", res3)
}
