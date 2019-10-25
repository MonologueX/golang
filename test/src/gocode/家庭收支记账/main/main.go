package main
import (
    "fmt"
)

func Menu() {
    fmt.Printf("--------------家庭收支记账软件--------------\n")
    fmt.Printf("\t\t1 收支明细\n")
    fmt.Printf("\t\t2 登记收入\n")
    fmt.Printf("\t\t3 登记支出\n")
    fmt.Printf("\t\t4 退    出\n")
    fmt.Printf("请选择(1-4): ")
}

func main() {
    // 初始化金额
    balance := 0.0
    // 收支记录
    details := "收支\t账户金额\t收支金额\t说    明"
    // 标记是否退出
    loop := true
    // 收支金额
    money := 0.0
    // 选择
    var choice int
    // 说明
    note := ""
    // 标记是否有收支
    flag := false
    for {
        Menu()
        // 输入选择
        fmt.Scanf("%d", &choice)
        switch choice {
            case 1:
                if flag {
                    fmt.Println("--------------当前收支明细记录--------------")
                    fmt.Println(details)
                } else {
                    fmt.Println("当前没有收支记录!")
                }
            case 2:
                fmt.Printf("请输入收入金额:")
                fmt.Scanf("%f", &money)
                balance += money
                fmt.Printf("请输入本次收入说明:")
                fmt.Scanf("%s", &note)
                details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
                flag = true
            case 3:
                fmt.Printf("请输入支出金额:")
                fmt.Scanf("%f", &money)
                if money > balance {
                    fmt.Println("余额不足!")
                } else {
                    balance -= money
                    fmt.Printf("请输入本次支出说明:")
                    fmt.Scanf("%s", &note)
                    details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
                    flag = true
                }
            case 4:
                var temp byte
                fmt.Printf("你确定要退出吗?(y/n):")
                fmt.Scanf("%c", &temp)
                if temp == 'y' {
                    loop = false
                }
            default:
                fmt.Printf("输入有误，请重新输入:")
        }
        if !loop {
            break
        }
        fmt.Println()
    }
    fmt.Println("你已经退出!")
}
