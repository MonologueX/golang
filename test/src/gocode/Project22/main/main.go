package main

import (
    "fmt"
    "sort"
    "math/rand"
)

type Hero struct {
    Name string
    Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
    return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
    return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
    temp := hs[i]
    hs[i] = hs[j]
    hs[j] = temp
}


func main() {
    var intSlice = []int{4, 1, 5, 6, 3, 8, 7, 9, 2}
    sort.Ints(intSlice)
    fmt.Println("排序后:", intSlice)

    var heroes HeroSlice
    for i := 0; i < 10; i++ {
        hero := Hero {
            Name:fmt.Sprintf("英雄~%d的", rand.Intn(100)),
            Age:rand.Intn(100),
        }
        heroes = append(heroes, hero)
    }
    fmt.Println("排序前:")
    for _, v := range heroes {
        fmt.Println(v)
    }
    sort.Sort(heroes)
    fmt.Println("排序后:")
    for _, v := range heroes {
        fmt.Println(v)
    }
}
