package main

import (
    "fmt"
    "strings"
)

// 闭包
// 累加器
func addUpter() (func(int) int) {
    var count int = 1
    return func(x int) int {
        count += x
        return count
    }
}

func makeSuffix(suffix string) func(string) string {
    return func(name string) string {
        if strings.HasSuffix(name, suffix) == false {
            return name + suffix
        }
        return name
    }
}

func main() {
    f2 := makeSuffix(".jpg")
    fmt.Printf("处理后: %q\n", f2("Make"))
    fmt.Printf("处理后: %q\n", f2("刘亦菲"))
    f := addUpter()
    for i := 0; i < 5; i++ {
        // 1 2 4 7 11
        fmt.Println(f(i))
    }
}
