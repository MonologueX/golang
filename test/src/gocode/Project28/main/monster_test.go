package monster
import (
    "testing"
)
func TestStore(t *testing.T) {
    monster := &Monster {
        Name:"Mary",
        Age:23,
        Skill:"c++",
    }
    res := monster.Store()
    if !res {
        t.Fatalf("monster.Store()错误，希望为%v,实际为%v", true, res)
    }
    t.Logf("monster.Store()正确!")
}
func TestRestore(t *testing.T) {
    var monster = &Monster{}
    res := monster.Restore()
    if !res {
        t.Fatalf("monster.Store()错误，希望为%v,实际为%v", true, res)
    }
    if monster.Name != "Mary" {
        t.Fatalf("monster.Store()错误，希望为%v,实际为%v", "Mary", monster.Name)
    }
    t.Logf("monster.Store()正确!")
}
