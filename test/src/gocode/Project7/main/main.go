package main
import "fmt"

func if_test() {
    var age int
    fmt.Printf("请输入年龄: ")
    fmt.Scanf("%d", &age)
    if age < 18 {
        fmt.Println("未成年!")
    } else if age < 50 {
        fmt.Println("青年!")
    } else {
        fmt.Println("老年!")
    }
}

func switch_test() {
    var num1 int
    fmt.Printf("请输入你要查询的星期: ")
    fmt.Scanf("%d", &num1)
    switch num1 {
        case 1:
            fmt.Println("计算机组成原理")
        case 2:
            fmt.Println("操作系统")
        case 3:
            fmt.Println("数据结构")
        case 4:
            fmt.Println("Java")
        case 5:
            fmt.Println("golang")
        case 6:
            fmt.Println("计算机网络")
        case 7:
            fmt.Println("C++")
        default:
            fmt.Println("输入不合法!")
    }
}

func main() {
    switch_test()
    //if_test()
}
