package main
import (
    "fmt"
    // 别名
    add "../math"
)

func fib(num int) int {
    if num <= 2 {
        return 1
    }
    return fib(num-1) + fib(num-2)
}

func swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func sum(a int, b int) (sum int) {
    sum = a + b
    return
}
type myFun func(int ,int) int
func mySum(cal myFun, a int, b int) int {
    return cal(a, b)
}

func main() {
    var a int = 6
    var b int = 2
    var c int
    c = add.Add(a, b)
    fmt.Println(c)
    c = fib(c)
    fmt.Println(c)
    fmt.Printf("a=%v,b=%v\n", a, b)
    swap(&a, &b)
    fmt.Printf("a=%v,b=%v\n", a, b)
    fmt.Printf("sum = %v\n", mySum(sum, a, b))
}
