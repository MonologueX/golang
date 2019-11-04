package main
import (
    "testing"
)
func TestGetSub(t *testing.T) {
    res := getSub(10, 3)
    if res != 7 {
        t.Fatalf("GetSub(10, 3)执行错误，期望值=%v, 实际值=%v\n", 7, res)
    }
    t.Logf("GetSub(10, 3) 执行正确!")
}
