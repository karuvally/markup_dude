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
    // Path to file
    var path = "sample.md"

    // Open the file in read mode
    var file, err = os.Open(path)

    // Produce error, if the file cannot be read
    if err != nil {
        log.Fatal(err)
    }
}

