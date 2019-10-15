package main
import (
    "fmt"
    "../test"
)

func init() {
    Test()
    fmt.Println("main::init")
}

func Test() {
    fmt.Println("main::test()")
}

func main() {
    fmt.Printf("age=%v, name=%q\n", test.Age, test.Name)
    fmt.Println("main()")
}
