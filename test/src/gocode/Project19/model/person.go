package model
import (
    "fmt"
)
type person struct {
    M_name string
    m_age int
    m_sal float64
}
func NewPerson(name string) *person {
    return &person{
        M_name: name,
    }
}

// age
func (p *person) SetAge(age int) {
    if age > 0 && age < 150 {
        p.m_age = age
    } else {
        fmt.Println("年龄有误！")
        p.m_age = 0
    }
}
func (p person)GetAge() int {
    return p.m_age
}

// sal
func (p *person) SetSal(sal float64) {
    if sal >= 3000 && sal <= 30000 {
        p.m_sal = sal
    } else {
        fmt.Println("薪水有误！")
        p.m_sal = 0
    }
}
func (p person)GetSal() float64 {
    return p.m_sal
}
