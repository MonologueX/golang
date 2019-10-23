package main

import (
    "fmt"
)

type Student struct {
    name string
    sex string
    age int
    phone string
}

func main() {
    var s1 Student
    s1.name = "Mark"
    s1.sex = "男"
    s1.age = 23
    s1.phone = "15892345234"

    fmt.Println("s1 =", s1)
    s2 := Student{"Mary", "女", 24, "13422435684"}
    fmt.Println("s2 =", s2)

    var s3 *Student = new(Student)
    (*s3).age = 25
    (*s3).name = "李朵"
    (*s3).phone = "15390438954"
    (*s3).sex = "女"
    fmt.Println("s3 =", *s3)

    var s4 *Student = &Student{}
    fmt.Println("s4 =", *s4)
}

// func main() {
//     arr := [5]int{3:-1}
//     fmt.Println("arr =", arr)
// }

// import (
//     "fmt"
//     "os"
// )
// 
// func main() {
//     var num int = 10
//     num += 10
//     fmt.Println("num =", num)
//     var s, sep string
//     for i := 1; i < len(os.Args); i++ {
//         s += sep + os.Args[i]
//         sep = " "
//     }
//     fmt.Println(s)
// }
