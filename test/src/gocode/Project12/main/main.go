package main

import (
    "fmt"
)

func sum(n1 int, n2 int) int {
    // 函数执行完毕后执行defer,栈
    defer fmt.Println("n1=", n1)
    defer fmt.Println("n2=", n2)
    n1++
    n2++
    res := n1 + n2
    fmt.Println("res=", res)
    return res
}

func main() {
    res1 := sum(1, 2)
    fmt.Println("res1=", res1)
    res2 := sum(3, 4)
    fmt.Println("res2=", res2)
    res3 := sum(5, 4)
    fmt.Println("res3=", res3)
}
