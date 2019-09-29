package main
import "fmt"

func fib(num int)(int) {
    if num <= 1 {
        return 1
    }
    return fib(num-1) + fib(num-2)
}

func main() {
    num := 5
    result := fib(num)
    fmt.Println(result)
}
