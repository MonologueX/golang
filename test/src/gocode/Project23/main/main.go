package main
import (
    "fmt"
)

type Usb interface {
    Start()
    Stop()
}

type Phone struct {
    Name string
}
func (p Phone) Start() {
    fmt.Println("phone start!")
}
func (p Phone) Stop() {
    fmt.Println("phone stop!")
}

type Camera struct {
    Name string
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
    var usbArr [3]Usb
    usbArr[0] = Phone{"vivo"}
    usbArr[1] = Phone{"小米"}
    usbArr[2] = Camera{"尼康"}
    fmt.Println(usbArr)
    computer := Computer{}
    phone := Phone{"小米"}
    camera := Camera{"尼康"}
    computer.Working(phone)
    computer.Working(camera)
}

