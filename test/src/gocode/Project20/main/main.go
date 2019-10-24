package main

import (
    "fmt"
)

// Student
type Student struct {
    Name string
    Age int
    Score int
}
func (stu Student) ShowInfo() {
    fmt.Printf("姓名:%q 年龄:%v 成绩:%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
    stu.Score = score
}

// Pupil
type Pupil struct {
    Student
}
func (p *Pupil) test() {
    fmt.Println("The pupil is testing ...")
}

// Graduate
type Graduate struct {
    Student
}
func (g *Graduate) test() {
    fmt.Println("The graduate is testing ...")
}

func main() {
    // pupil
    pupil := &Pupil{}
    pupil.Student.Name = "Mark"
    pupil.Student.Age = 12
    pupil.test()
    pupil.Student.SetScore(80)
    pupil.Student.ShowInfo()

    // Graduate
    graduate := &Graduate{}
    graduate.Student.Name = "Tom"
    graduate.Student.Age = 23
    graduate.test()
    graduate.Student.SetScore(90)
    graduate.Student.ShowInfo()
}
