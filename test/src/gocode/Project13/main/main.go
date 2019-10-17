package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    str := "hello go 你好"
    fmt.Printf("len = %v\n", len(str))
    // 遍历不带中文
    for i := 0; i < len(str); i++ {
        fmt.Printf("%c ", str[i])
    }
    fmt.Println()
    // 遍历中文
    r := []rune(str)
    for i := 0; i < len(r); i++ {
        fmt.Printf("%c ", r[i])
    }
    fmt.Println()
    // 字符串转整数
    n, err := strconv.Atoi("123")
    // n, err := strconv.Atoi("123hello")
    if err != nil {
        fmt.Println("转换错误", err)
    } else {
        fmt.Println("转换正确", err)
        fmt.Println("num =", n)
    }
    // 整数转字符串
    fmt.Printf("整数转字符串:%q\n", strconv.Itoa(123))

    // 字符串转byte
    bytes := []byte("hello go")
    fmt.Printf("bytes=%v\n", bytes)

    // byte转字符串
    str3 := string([]byte{97, 98, 99})
    fmt.Printf("str3=%q\n", str3)

    // 查找字符串
    b := strings.Contains("hasdgf", "dghj")
    if !b {
        fmt.Println("没有找到:", b)
    } else {
        fmt.Println("找到了:", b)
    }

    // 统计字符个数
    fmt.Printf("num=%v\n", strings.Count("afdgfghdklhg", "g"))

    // 字符串比较
    // 不区分大小写
    b = strings.EqualFold("abcdE", "abcde")
    if !b {
        fmt.Println("不相等:", b)
    } else {
        fmt.Println("相等:", b)
    }

    // 查找下标
    fmt.Printf("index=%v\n", strings.Index("asdfgh", "fgh"))
    fmt.Printf("LastIndex=%v\n", strings.LastIndex("fghasdfgh", "fgh"))

    // 替换
    fmt.Println("str =", strings.Replace("hello go go", "go go", "go", 1))

    // 分割
    strArray := strings.Split("hello,go,xiao,ying",",")
    for i := 0; i < len(strArray); i++ {
        fmt.Printf("%v ", strArray[i])
    }
    fmt.Println()

    // 大小写转换
    fmt.Println("str =", strings.ToLower("Go"))
    fmt.Println("str =", strings.ToUpper("Go"))

    // 去掉左右两边空格
    // strings.TrimLeft or TrimRight
    fmt.Println("str =", strings.TrimSpace(" hello go "))

    // 去掉指定字符
    fmt.Println("str =", strings.Trim(",,,hello go,,,", ","))

    // strings.HasSuffix
    // strings.HasPrefix
}
