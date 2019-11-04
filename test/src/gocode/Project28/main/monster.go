package monster

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
)

type Monster struct {
    Name string
    Age int
    Skill string
}

func (this *Monster) Store() bool {
    data, err := json.Marshal(this)
    if err != nil {
        fmt.Println("Monster err =", err)
        return false
    }
    path := "monster.ser"
    err = ioutil.WriteFile(path, data, 0666)
    if err != nil {
        fmt.Println("Write file =", err)
        return false
    }
    return true
}

func (this *Monster) Restore() bool {
    path := "monster.ser"
    data, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("Read file =", err)
        return false
    }
    err = json.Unmarshal(data, this)
    if err != nil {
        fmt.Println("Unmarshal err =", err)
        return false
    }
    return true
}
