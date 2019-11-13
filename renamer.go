package main

import (
    "fmt"
    "strings"
    "os"
)

func main() {

    fmt.Println(os.Args[1:3])
    // Example 1: Willkommen to GoLangCode.com
    myText := string(os.Args[1:2])
    myText = strings.Replace(myText, "[pseudo] ", "", -1)
    fmt.Println(myText)

    // Example 2: Change first occurance
    // Output: The car sounds sound
    myText = "The sound sounds sound"
    myText = strings.Replace(myText, "sound", "car", 1)
    fmt.Println(myText)

    // Example 3: Replacing quotes (double backslash needed)
    // Output: I \'quote\' this text
    myText = "I 'quote' this text"
    myText = strings.Replace(myText, "'", "\\'", -1)
    fmt.Println(myText)
}