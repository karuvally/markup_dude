// markup_dude, Golang edition!

// Import required libraries
import (
    "os"
    "fmt"
)

// The main function
func main() {
    var path = "sample.md"
    var file, err = os.Open(path)

    if err != nil {
        log.Fatal(err)
    }
}

