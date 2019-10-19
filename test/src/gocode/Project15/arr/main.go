package main

import (
    "fmt"
)

func testArr() {
    // 初始化
    var arr1[3]int = [3]int{1, 3, 5}
    fmt.Println("arr1 =", arr1)

    var arr2 = [3]int{1, 3, 5}
    fmt.Println("arr2 =", arr2)

    var arr3 = [...]int{1, 3, 5}
    fmt.Println("arr3 =", arr3)

    var arr4 = [...]int{2:1, 1:3, 0:5}
    fmt.Println("arr4 =", arr4)

    arr5 := [...]string{2:"make", 1:"hello", 0:"hi"}
    fmt.Println("arr5 =", arr5)

    arr6 := [...]int{2:1, 1:3, 0:5}
    fmt.Println("arr3 =", arr6)
}

func main() {
    testArr()
    var arr[5]int
    for i := 0; i < len(arr); i++ {
        fmt.Scanf("%d", &arr[i])
    }
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%v ", arr[i])
    }
    fmt.Println()
    for index, value := range arr {
        fmt.Printf("%v:%v ", index, value)
    }
    fmt.Println()
}
