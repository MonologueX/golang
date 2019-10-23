package main
import (
    "fmt"
)

type Date struct {
    m_day int
    m_month int
    m_year int
    m_hour int
    m_minutes int
    m_second int
}

func (d *Date) dateSet(year int, month int, day int) {
    d.m_year = year
    d.m_month = month
    d.m_day = day
}

func (d Date) datePrint() {
    fmt.Printf("%v-%v-%v\n", d.m_year, d.m_month, d.m_day)
}

func (d *Date) timeSet(hour int, minutes int, second int) {
    d.m_hour = hour
    d.m_minutes = minutes
    d.m_second = second
}

func (d Date) timePrint() {
    fmt.Printf("%v:%2.v\n", d.m_hour, d.m_minutes)
}

func main() {
    var d Date
    d.timeSet(17, 2, 4)
    d.dateSet(2019, 10, 23)
    d.datePrint()
    d.timePrint()
}
