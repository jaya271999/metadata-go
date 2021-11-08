package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    matches := filepath.Glob("/proc/*/exe")
    for _, file := range matches {
        target, _ := os.Readlink(file)
        if len(target) > 0 {
            fmt.Printf("%+v\n", target)
        }
    }
}
