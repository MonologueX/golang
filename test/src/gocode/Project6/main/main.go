package main
import (
    "fmt"
    "strconv"
)

// 基本类型转字符串
func to_string_test() {
    var num1 int = 99
    var num2 float64 = 23.456
    var b bool = true
    var my_char byte = 'h'
    var str string

    // Sprintf
    str = fmt.Sprintf("%d", num1)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%f", num2)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%t", b)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%c", my_char)
    fmt.Printf("str type %T str=%q\n\n", str, str)

    var num3 int = 66
    var num4 float64 = 3.1415
    var b2 bool = false

    str = strconv.FormatInt(int64(num3), 10)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = strconv.FormatFloat(num4, 'f', 10, 64)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = strconv.FormatBool(b2)
    fmt.Printf("str type %T str=%q\n\n", str, str)
}

func string_to_basictype() {
    var str1 string = "true"
    var b bool
    b, _ = strconv.ParseBool(str1)
    fmt.Printf("b type %T b=%v\n", b, b)

    var str2 string = "123456"
    var num1 int64
    num1, _ = strconv.ParseInt(str2,10, 64)
    fmt.Printf("num1 type %T num1=%v\n", num1, num1)

    var str3 string = "3.1415"
    var f1 float64
    f1, _ = strconv.ParseFloat(str3, 64)
    fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}

func main() {
    to_string_test()
    string_to_basictype()
}
