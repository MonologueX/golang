package main

import(
    "fmt"
    "../model"
)

func main() {
    // stu := model.Student{
    //     Name:"Mark",
    //     Score:90,
    // }
    stu := model.NewStudent("Mark", 90)
    fmt.Println(*stu)
}
