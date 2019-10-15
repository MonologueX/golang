package test
import "fmt"

var Age int
var Name string

func init() {
    Age = 20
    Name = "Mark"
    fmt.Println("test::init()")
}
