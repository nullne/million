// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {


    // For more granular writes, open a file for writing.
    f, err := os.Create("/tmp/dat2")
    check(err)

    // It's idiomatic to defer a `Close` immediately
    // after opening a file.
    defer f.Close()


    // A `WriteString` is also available.
    n3, err := f.WriteString(fmt.Sprintf("%-13s%-20s%20s\n", "", "you", "me"))
    fmt.Printf("wrote %d bytes\n", n3)

    // Issue a `Sync` to flush writes to stable storage.
    f.Sync()


}

