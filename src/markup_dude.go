// markup_dude, Golang edition!

// This is the main package
package main

// Import required libraries
import (
    "os"
    "log"
)

// The main function
func main() {
    var path = "sample.md"
    var file, err = os.Open(path)

    if err != nil {
        log.Fatal(err)
    }
}

