package main

import (
    "github.com/edervs/Modules-in-Go/submodule"
    "github.com/edervs/Modules-in-Go/submodule2"

    submodulev2 "github.com/edervs/Modules-in-Go/submodule/v2"
    submodulev3 "github.com/edervs/Modules-in-Go/submodule/v3"
)

func main() {
    submodule.PrintSubmodule()
    submodulev2.PrintSubmodule("Person")
    submodulev3.PrintSubmodule("Person", 18)
    submodule2.Welcome("Person")
}

