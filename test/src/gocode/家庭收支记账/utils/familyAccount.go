package utils
import (
    "fmt"
)

type familyAccount struct {
    // 初始化金额
    balance float64
    // 收支记录
    details string
    // 标记是否退出
    loop bool
    // 收支金额
    money float64
    // 说明
    note string
    // 标记是否有收支
    flag bool
}

func NewFamilyAccount() *familyAccount {
    return &familyAccount{
        balance : 0.0,
        details : "收支\t账户金额\t收支金额\t说    明",
        loop : true,
        money : 0.0,
        note : "",
        flag : false,
    }
}

// 主菜单
func (this familyAccount) Menu() {
    fmt.Printf("--------------家庭收支记账软件--------------\n")
    fmt.Printf("\t\t1 收支明细\n")
    fmt.Printf("\t\t2 登记收入\n")
    fmt.Printf("\t\t3 登记支出\n")
    fmt.Printf("\t\t4 退    出\n")
    fmt.Printf("请选择(1-4): ")
}

// 收支记录
func (this familyAccount) Show() {
    if this.flag {
        fmt.Println("--------------当前收支明细记录--------------")
        fmt.Println(this.details)
    } else {
        fmt.Println("当前没有收支记录!")
    }
}

// 收入记录
func (this *familyAccount) Income() {
    fmt.Printf("请输入收入金额:")
    fmt.Scanf("%f", &this.money)
    this.balance += this.money
    fmt.Printf("请输入本次收入说明:")
    fmt.Scanf("%s", &this.note)
    this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
    this.flag = true
}

// 支出记录
func (this *familyAccount) Pay() {
    fmt.Printf("请输入支出金额:")
    fmt.Scanf("%f", &this.money)
    if this.money > this.balance {
        fmt.Println("余额不足!")
    } else {
        this.balance -= this.money
        fmt.Printf("请输入本次支出说明:")
        fmt.Scanf("%s", &this.note)
        this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
        this.flag = true
    }
}

// 退出
func (this *familyAccount) Exit() {
    var temp byte
    fmt.Printf("你确定要退出吗?(y/n):")
    fmt.Scanf("%c", &temp)
    if temp == 'y' {
        this.loop = false
    }
}

func (this *familyAccount) Start () {
    for {
        this.Menu()
        // 输入选择
        var choice int
        fmt.Scanf("%d", &choice)
        switch choice {
            case 1:
                this.Show()
            case 2:
                this.Income()
            case 3:
                this.Pay()
            case 4:
                this.Exit()
            default:
                fmt.Printf("输入有误，请重新输入:")
        }
        if !this.loop {
            break
        }
        fmt.Println()
    }
    fmt.Println("你已经退出!")
}
