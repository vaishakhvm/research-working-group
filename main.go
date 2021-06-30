package main

import (
    "fmt"
    "testing"
)

func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {

        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

func helloworld() string {
	return "Hello   World!!!"
}

func main() {
	fmt.Println(helloworld())
	fmt.Println("This a lucky number to be ignored:",IntMin(10,20))
}
