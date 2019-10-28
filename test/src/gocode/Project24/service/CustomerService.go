package service

import (
    "../model"
    "fmt"
)

type CustomerService struct {
    customers []model.Customer
    num int
}

func NewCustomerService() *CustomerService {
    cs := &CustomerService{}
    cs.num = 1
    c := model.NewCustomer(1, "许玲珑", "女", 23, "15892345690", "xll199025@163.com")
    cs.customers = append(cs.customers, c)
    return cs
}

func (this *CustomerService) List()[]model.Customer {
    return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
    this.num++
    customer.Id = this.num
    this.customers = append(this.customers, customer)
    return true;
}

func (this CustomerService) FindById(id int) int {
    index := -1
    for i := 0; i < len(this.customers); i++ {
        if this.customers[i].Id == id {
            index = i
            break
        }
    }
    return index
}

func (this *CustomerService) Delete(id int) bool {
    index := this.FindById(id)
    if index == -1 {
        return false
    }
    this.customers = append(this.customers[:index], this.customers[index+1:]...)
    return true
}

func (this *CustomerService) Modify(id int) bool {
    index := this.FindById(id)
    if index == -1 {
        return false
    }
    fmt.Printf("姓名:")
    fmt.Scanln(&this.customers[index].Name)
    fmt.Printf("性别:")
    fmt.Scanln(&this.customers[index].Gender)
    fmt.Printf("年龄:")
    fmt.Scanln(&this.customers[index].Age)
    fmt.Printf("电话:")
    fmt.Scanln(&this.customers[index].Phone)
    fmt.Printf("电子邮件:")
    fmt.Scanln(&this.customers[index].Email)
    return true
}
