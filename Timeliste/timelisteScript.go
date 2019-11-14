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
  EarliestTime := "08"
  LatestTime := "16"

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
  workingDays := make([]string, 0)

  for n := range words {
    // check to retrieve current logged in users activity
    if(words[n] == user[0]){
      // retrieving values from current month
      if(words[n+4] == monthShort){
        date := words[n+5]
        endTime := ""

        if(words[n+7] != "-" || words[n+8] == "down"){ // ignoring when user is still logged in.
          // endTime = words[n+7] + " " + words[n+8]
          continue
        } else{
          endTime = words[n+8]
        }
        
        startTime := words[n+6] // setting start time of workshift
        
        // time := startTime + " - " + endTime
        // fmt.Println(date, monthString, time)


        // checking if user logged in after worktime started or before
        if(startTime[0:2] >= EarliestTime) {
          startTime = startTime[0:2] + ":00"
        } else {
          startTime = EarliestTime + ":00"
        }

        // checking if user was still logged in when shift ended
        if(endTime[0:2] <= LatestTime) {
          endTime = endTime[0:2] + ":00"
        } else {
          endTime = LatestTime + ":00"
        }

        workTime := startTime + " - " + endTime

        workingDays = append(workingDays, (date + " " + monthString + " " + workTime))
      }
    }
  }

  workingDays = removeDuplicates(workingDays)

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