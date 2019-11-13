package main

import (
    "fmt"
    "strings"
    "os"
    "log"
)

func main() {
    // function that renames the files in current directory
    // goes trough all the files, searches for the words after "file.Name()" in the replace method
    // and replaces those words with whatever is in the next portion.

    dirname := "."

    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        myText := strings.Replace(file.Name(), " [1080p] [h.265]", "", -1)
        fmt.Println(myText)
        //fmt.Println(file.Name())
        os.Rename(file.Name(), myText)
        f.WriteString(myText)
    }

    f.Close()
}