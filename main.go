package main

import (
    "github.com/edervs/Modules-in-Go/submodule"

    submodulev2 "github.com/edervs/Modules-in-Go/submodule/v2"
)

func main() {
    submodule.PrintSubmodule()
    submodulev2.PrintSubmodule("Person")
}

