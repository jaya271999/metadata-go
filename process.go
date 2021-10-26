package main

import (
    "fmt"
    "os"
    "filepath"
)

func main() {
    matches, err := filepath.Glob("/proc/*/exe")
    for _, file := range matches {
        target, _ := os.Readlink(file)
        if len(target) > 0 {
            fmt.Printf("%+v\n",target)
        }
    }
}
