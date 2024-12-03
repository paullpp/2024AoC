package main

import (
  "fmt"
  "strings"
  "os"
  "bufio"
  "strconv"
)

func handle_err(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  // get data from file
  file, err := os.Open("inputs.txt")
  handle_err(err)

  defer file.Close()

  // scan line-by-line and add values to lists
  scanner := bufio.NewScanner(file)

  safe := 0

  for scanner.Scan() {
    line := scanner.Text()
    parsed_line := strings.Split(line, " ")
  
    var decreasing bool

    parsed_first, err := strconv.Atoi(parsed_line[0])
    handle_err(err)

    parsed_second, err := strconv.Atoi(parsed_line[1])
    handle_err(err)

    if parsed_first < parsed_second {
      decreasing = false
    } else {
      decreasing = true
    }

    for i := 0; i < len(parsed_line) - 1; i++ {
      parsed_curr, err := strconv.Atoi(parsed_line[i])
      handle_err(err)

      parsed_next, err := strconv.Atoi(parsed_line[i+1])
      handle_err(err)

      if decreasing {
        val := parsed_curr - parsed_next
        if val < 1 || val > 3 {
          safe -= 1
          break
        }
      } else {
        val := parsed_next - parsed_curr
        if val < 1 || val > 3 {
          safe -= 1
          break
        }  
      }
    }

    safe += 1
  }
  
  // part 1
  fmt.Println("===", safe, "safe reports")

  // TODO: part 2
}

