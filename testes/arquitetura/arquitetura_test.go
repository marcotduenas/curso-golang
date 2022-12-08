package arquitetura

import(
    "testing"
    "runtime"
)

func TestDependent(t * testing.T){
    if runtime.GOARCH == "amd64"{
        t.Skip("Nao funciona em arquitetura amd64")
    }

    //...
    t.Fail()
}

