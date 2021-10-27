package main

import (
    "fmt"
    "github.com/pbnjay/memory"
)

func main() {
    fmt.Printf("Total system memory: %d\n", memory.TotalMemory())
}
