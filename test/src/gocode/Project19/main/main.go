/***********************************
 文件名称: main.go
 文件描述: 封装练习
 编译环境: Linux
 作者相关: 心文花雨
***********************************/

package main
import (
    "fmt"
    "../model"
)

func main() {
    p := model.NewPerson("Mark")
    p.SetAge(23)
    p.SetSal(4000)
    fmt.Printf("name=%q, age=%v, sal=%v\n", p.M_name, p.GetAge(), p.GetSal())
    fmt.Println("p =", *p)
}
