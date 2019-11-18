package main

import (
    "github.com/gomodule/redigo/redis"
    "fmt"
)

func main() {
    conn, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("redis error!", err)
        return
    }

    _,err = conn.Do("AUTH","123456")
    if err != nil {
        fmt.Println("auth failed:", err)
    }

    defer conn.Close()

    _, err = conn.Do("Set", "s1", "hello golang")
    if err != nil {
        fmt.Println("set error!", err)
        return
    }

    r, err := redis.String(conn.Do("Get", "s1"))
    if err != nil {
        fmt.Println("get error!", err)
        return
    }

    fmt.Println("s1 =", r)
    fmt.Println("ok!")
}
