package main
import (
    "fmt"
    "encoding/json"
)

type Monster struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Skill string `json:"skill"`
}

func main() {
    monster :=Monster{"孙悟空", 1000, "金箍棒"}
    fmt.Println("m =", monster)
    jsonStr, err := json.Marshal(monster)
    if err != nil {
        fmt.Println("json 处错误", err)
    }
    fmt.Println("jsonStr=", string(jsonStr))
}
