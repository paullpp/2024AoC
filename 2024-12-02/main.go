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

func is_safe(list []string) bool {
  var decreasing bool

  if len(list) < 2 {
    return true
  }

  parsed_first, err := strconv.Atoi(list[0])
  handle_err(err)

  parsed_second, err := strconv.Atoi(list[1])
  handle_err(err)

  if parsed_first < parsed_second {
    decreasing = false
  } else {
    decreasing = true
  }

  for i := 0; i < len(list) - 1; i++ {
    parsed_curr, err := strconv.Atoi(list[i])
    handle_err(err)

    parsed_next, err := strconv.Atoi(list[i+1])
    handle_err(err)

    if decreasing {
      val := parsed_curr - parsed_next
      if val < 1 || val > 3 {
        return false
      }
    } else {
      val := parsed_next - parsed_curr
      if val < 1 || val > 3 {
        return false
      }  
    }
  }

  return true
}

func main() {
  file, err := os.Open("inputs.txt")
  handle_err(err)

  defer file.Close()

  scanner := bufio.NewScanner(file)

  safe := 0
  safe_dampened := 0

  for scanner.Scan() {
    line := scanner.Text()
    parsed_line := strings.Split(line, " ")

    if is_safe(parsed_line) {
      safe += 1
    } else {
      for i := 0; i < len(parsed_line); i++ {
        new_list := append([]string(nil), parsed_line[:i]...)
        new_list = append(new_list, parsed_line[i+1:]...)

        if is_safe(new_list) {
            safe_dampened += 1
            break
        } else {
        }
      }
    }
  }
  
  // part 1
  fmt.Println("===", safe, "safe reports")

  // part 2
  safe_dampened += safe
  fmt.Println("===", safe_dampened, "safe dampened reports")
}

