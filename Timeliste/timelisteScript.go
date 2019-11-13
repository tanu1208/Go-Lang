package main

import (
  "fmt"
  "os/exec"
  "runtime"
  "time"
  "strings"
)

func execute() {
  monthTime := time.Now().Month()
  monthString := monthTime.String()
  monthShort := monthString[0:3]

  out, err := exec.Command("whoami").Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  output := string(out)
  user := strings.Fields(output)

  out2, err2 := exec.Command("last").Output()
  if err2 != nil {
    fmt.Printf("%s", err2)
  }

  output2 := string(out2)
  words := strings.Fields(output2)
  dates := make([]string, 0)

  if(words[0] == user[0]){
    for n := range words {
      if(words[n] == monthShort){
        // fmt.Println(words[n+1] + " " + monthString)
        tmp := string(words[n+2])
        time := string(tmp[0:2]) + ":00"
        dates = append(dates, (words[n+1] + " " + monthString + " " + time))
      }
    }
  }

  workingDays := removeDuplicates(dates)

  fmt.Println(workingDays)
}

func removeDuplicates(elements []string) []string {
  // Use map to record duplicates as we find them.
  encountered := map[string]bool{}
  result := []string{}

  for v := range elements {
    if encountered[elements[v]] == true {
      // Do not add duplicate.
    } else {
      // Record this element as an encountered element.
      encountered[elements[v]] = true
      // Append to result slice.
      result = append(result, elements[v])
    }
  }
  // Return the new slice.
  return result
}

func main() {
  if runtime.GOOS == "windows" {
    fmt.Println("Can't Execute this on a windows machine")
  } else {
    execute()
  }
}