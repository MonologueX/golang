package main
import (
    "fmt"
)

type Usb interface {
    Start()
    Stop()
}

type Phone struct {
}
func (p Phone) Start() {
    fmt.Println("phone start!")
}
func (p Phone) Stop() {
    fmt.Println("phone stop!")
}

type Camera struct {
}
func (c Camera) Start() {
    fmt.Println("Camera start!")
}
func (c Camera) Stop() {
    fmt.Println("Camera stop!")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
    usb.Start()
    usb.Stop()
}

func main() {
    computer := Computer{}
    phone := Phone{}
    camera := Camera{}
    computer.Working(phone)
    computer.Working(camera)
}
