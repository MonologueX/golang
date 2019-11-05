package main

import (
    "fmt"
)

func main() {
    // 引用类型
    var intChan chan int
    intChan = make(chan int, 3)
    fmt.Printf("intChan 的值是%v, 本省地址是%v\n", intChan, &intChan)
    // 写入数据
    intChan <-10
    num := 200
    intChan <-num
    // 查看容量和长度
    fmt.Printf("len=%v, cap=%v\n", len(intChan), cap(intChan))
    // 读取数据
    var num2 int
    num2 = <-intChan
    fmt.Printf("len=%v, cap=%v\n", len(intChan), cap(intChan))
    fmt.Println(num2)
    // num3 := <-intChan
    // num4 := <-intChan
}
