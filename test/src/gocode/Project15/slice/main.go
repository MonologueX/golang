package main

import (
    "fmt"
)

func main() {
    // 第一种方式
    var intArr[5]int = [...]int{1, 2, 3, 4, 5}
    slice := intArr[1:3] // 1 <= index < 3
    fmt.Println("slce =", slice)
    fmt.Println("len =", len(slice))
    fmt.Println("cap =", cap(slice))
    for i := 0; i < len(slice); i++ {
        fmt.Printf("slice[%v]=%v ", i, slice[i])
    }
    fmt.Println()
    fmt.Println()

    // 第二种方式
    var slice2 []int = make([]int, 4, 10)
    fmt.Println("slce =", slice2)
    fmt.Println("len =", len(slice2))
    fmt.Println("cap =", cap(slice2))
    fmt.Println()

    // 第三种方式
    var slice3 []string = []string{"Tom", "Make", "Mary"}
    fmt.Println("slce =", slice3)
    fmt.Println("len =", len(slice3))
    fmt.Println("cap =", cap(slice3))
    for i, v := range slice3 {
        fmt.Printf("slice[%v]=%v ", i, v)
    }
    fmt.Println()

    // append
    var slice4[]int = []int{1, 2, 3}
    slice4 = append(slice4, 4, 5, 6)
    slice4 = append(slice4, slice4 ...)
    fmt.Println("slice4 = ", slice4)

    // copy
    var slice5[]int = []int{1, 2, 3, 4, 5}
    var slice6 = make([]int, 10)
    copy(slice6, slice5)
    fmt.Println("slice6 =", slice6)
}
