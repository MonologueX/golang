package main
import (
    "fmt"
    "strconv"
)

func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    size := len(str)
    for i := 0; i < size; i++ {
        if str[i] == str[size-1] {
            size--
        } else {
            return false;
        }
    }
    fmt.Println(str)
    fmt.Println(size)
    return true
}

func main() {
    fmt.Println(isPalindrome(3))
    m := make(map[string]string)
    m["1"] = "1"
    m["2"] = "1"
    m["3"] = "1"
    fmt.Println(m)
}
