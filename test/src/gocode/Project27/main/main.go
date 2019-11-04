/***********************************
 文件名称: main.go
 文件描述: json
 编译环境: Linux
 作者相关: 心文花雨
***********************************/
package main
import (
    "fmt"
    "encoding/json"
)
type Monster struct {
    // 反射
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday string `json:"birthday"`
    Sal float64 `json:"sal"`
    Skill string `json:"skill"`
}

func testStruct() {
    monster := Monster {
        Name:"Mary",
        Age:24,
        Birthday:"1996-11-4",
        Sal:97,
        Skill:"golang",
    }
    // 序列化
    data, err := json.Marshal(&monster)
    if err != nil {
        fmt.Println("序列化失败,", err)
    }
    fmt.Println("monster序列化后:", string(data))
    // 反序列化
    var monster2 Monster
    err = json.Unmarshal([]byte(string(data)), &monster2)
    if err != nil {
        fmt.Println("反序列化失败,", err)
    }
    fmt.Println("反序列化后:", monster2)
}

func testMap() {
    var m map[string]interface{}
    m = make(map[string]interface{})
    m["name"] = "Bob"
    m["age"] = 23
    m["birthday"] = "1996-12-4"
    m["sal"] = 98
    m["skill"] = "c++"
    // 序列化
    data, err := json.Marshal(m)
    if err != nil {
        fmt.Println("序列化失败,", err)
    }
    fmt.Println("map序列化后:", string(data))
    // 反序列化
    var m2 map[string]interface{}
    err = json.Unmarshal([]byte(string(data)), &m2)
    if err != nil {
        fmt.Println("反序列化失败,", err)
    }
    fmt.Println("反序列化后:", m2)
}

func testSlice() {
    var slice []map[string]interface{}
    var m1 map[string]interface{}
    m1 = make(map[string]interface{})
    m1["name"] = "Bob"
    m1["age"] = 23
    m1["skill"] = "c++"
    slice = append(slice, m1)

    var m2 map[string]interface{}
    m2 = make(map[string]interface{})
    m2["name"] = "Mary"
    m2["age"] = 25
    m2["skill"] = "Java"
    slice = append(slice, m2)
    // 序列化
    data, err := json.Marshal(slice)
    if err != nil {
        fmt.Println("序列化失败,", err)
    }
    fmt.Println("slice序列化后:", string(data))
    // 反序列化
    var s []map[string]interface{}
    err = json.Unmarshal([]byte(string(data)), &s)
    if err != nil {
        fmt.Println("反序列化失败,", err)
    }
    fmt.Println("反序列化后:", s)
}
func main() {
    testStruct()
    testMap()
    testSlice()
}
