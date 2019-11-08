package main

import (
    "fmt"
    "reflect"
)

// 声明定义了结构体
type Monster struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Score float32 `json:"score"`
    Sex string `json:"sex"`
}

func (s Monster) Print() {
    fmt.Println("start!")
    fmt.Println(s)
    fmt.Println("end")
}

func (s Monster) GetSum(n1, n2 int) int {
    return n1 + n2
}

func (s* Monster) Set(name string, age int, score float32, sex string) {
    s.Name = name
    s.Age = age
    s.Score = score
    s.Sex = sex
}

func testStruct(a interface{}) {
    typ := reflect.TypeOf(a)
    value := reflect.ValueOf(a)
    kind := value.Kind()
    if kind != reflect.Struct {
        fmt.Println("expect struct!")
        return
    }
    num := value.NumField()
    fmt.Printf("struct has %d fields!\n", num)
    // 遍历结构体字段
    for i := 0; i < num; i++ {
        fmt.Printf("field %d:%v\n", i, value.Field(i))
        tagVal := typ.Field(i).Tag.Get("json")
        if tagVal != "" {
            fmt.Printf("filed %d tag:%v\n", i, tagVal)
        }
    }
    numOfMethod := value.NumMethod()
    fmt.Printf("struct has %d method!\n", numOfMethod)
    // 按照ASCII码排序
    value.Method(1).Call(nil)
    var param []reflect.Value
    param = append(param, reflect.ValueOf(10))
    param = append(param, reflect.ValueOf(40))
    res := value.Method(0).Call(param)
    fmt.Println("res =", res[0].Int())
}

func main() {
    var a Monster = Monster {
        Name : "Mary",
        Age : 23,
        Score : 90,
    }
    testStruct(a)
}
