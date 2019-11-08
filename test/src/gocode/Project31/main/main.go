package main
import (
    "fmt"
    "reflect"
)

func reflectTest(b interface{}) {
    rType := reflect.TypeOf(b)
    fmt.Println("rType =", rType)
    rValue := reflect.ValueOf(b)
    fmt.Printf("rValue=%v, type=%T\n", rValue, rValue)
    num := 2 + rValue.Int()
    fmt.Println("num =", num)
    iV := rValue.Interface()
    num2 := iV.(int)
    fmt.Println("num2 =", num2)
}

func main() {
    var num int = 100
    reflectTest(num)
}
