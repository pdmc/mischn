package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    f, _ := os.Open("file")

    r4 := bufio.NewReader(f)
    b4, _ := r4.Peek(5)

    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close()
}
