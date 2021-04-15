package submodule

import "fmt"

func PrintSubmodule(name string, age int) {
    fmt.Printf("Hello, %s. This is the submodule in the version %s\n", name, version)
    fmt.Printf("You are %d years old.\n", age)
}
