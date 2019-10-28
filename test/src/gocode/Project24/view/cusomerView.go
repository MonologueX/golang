package main
import (
    "fmt"
    "../service"
    "../model"
)

type customerView struct {
    key int
    loop bool
    cs *service.CustomerService
}

func (this customerView) menu() {
    fmt.Printf("=============主菜单=============\n")
    fmt.Printf("\t1 添加客户\n")
    fmt.Printf("\t2 修改客户\n")
    fmt.Printf("\t3 删除客户\n")
    fmt.Printf("\t4 客户列表\n")
    fmt.Printf("\t5 退 出\n")
    fmt.Print("请输入：")
}

func (this customerView) list() {
    customers := this.cs.List()
    fmt.Println("=============================客户列表=============================")
    fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
    for i := 0; i < len(customers); i++ {
        fmt.Println(customers[i].GetInfo())
    }
    fmt.Println()
}

func (this customerView) add() {
    fmt.Println("添加客户")
    fmt.Printf("姓名:")
    name := ""
    fmt.Scanln(&name)
    fmt.Printf("性别:")
    gender := ""
    fmt.Scanln(&gender)
    fmt.Printf("年龄:")
    age := 0
    fmt.Scanln(&age)
    fmt.Printf("电话:")
    phone := ""
    fmt.Scanln(&phone)
    fmt.Printf("电子邮件:")
    email := ""
    fmt.Scanln(&email)

    customer := model.NewCustomer2(name, gender, age, phone, email)
    if this.cs.Add(customer) {
        fmt.Println("添加成功!")
    } else {
        fmt.Println("添加失败!")
    }
    fmt.Println()
}

func (this *customerView) delete() {
    fmt.Printf("请输入要删除客户的编号，输入-1退出:")
    id := -1
    fmt.Scanf("%d\n", &id)
    if id == -1 {
        return
    }
    flag := "y"
    fmt.Printf("请确认是否删除(y/n):")
    fmt.Scanln(&flag)
    if flag == "n" {
        return
    }
    if this.cs.Delete(id) {
        fmt.Println("删除成功!")
    } else {
        fmt.Println("删除失败!")
    }
    fmt.Println()
    return
}

func (this *customerView) modify() {
    fmt.Printf("请输入要修改客户的编号，输入-1退出:")
    id := -1
    fmt.Scanf("%d\n", &id)
    if id == -1 {
        return
    }
    if this.cs.Modify(id) {
        fmt.Println("修改成功!")
    } else {
        fmt.Println("修改失败!")
    }
    fmt.Println()
    return
}

func (this *customerView) exit() {
    fmt.Printf("确认是否退出(y/n):")
    flag := 'y'
    fmt.Scanln(&flag)
    if flag == 'y' {
        this.loop = false
    }
}

func (this *customerView) Start() {
    for {
        this.menu()
        fmt.Scanf("%d\n", &this.key)
        switch this.key {
            case 1:
                this.add()
            case 2:
                this.modify()
            case 3:
                this.delete()
            case 4:
                this.list()
            case 5:
                this.exit()
            default:
                fmt.Println("你的输入有误，请重新输入!")
        }
        if !this.loop {
            break
        }
    }
    fmt.Println("你已经退出!")
}

func main() {
    customer := customerView {
        key : 0,
        loop : true,
    }
    customer.cs = service.NewCustomerService()
    customer.Start()
}
