package main

import (
  "fmt"
  "math"
  "slices"
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

func get_occurences(list []int, value int) int {
  count := 0
  for i := 0; i < len(list); i++ {
    if list[i] == value {
      count++
    }
  }

  return count
}

func main() {
  /*
    1. sort both lists in ascending order
    2. pop from both lists
    3. take the difference, add to output
    4. return output
  */

  // get data from file
  file, err := os.Open("inputs.txt")
  handle_err(err)

  defer file.Close()

  // scan line-by-line and add values to lists
  scanner := bufio.NewScanner(file)

  var left_list, right_list []int

  for scanner.Scan() {
    line := scanner.Text()
    parsed_line := strings.Split(line, "   ")
    
    if len(parsed_line) != 2 {
      panic("found more than two numbers in one line")
    }

    parsed_first, err := strconv.Atoi(parsed_line[0])
    handle_err(err)

    parsed_second, err := strconv.Atoi(parsed_line[1])
    handle_err(err)

    left_list = append(left_list, parsed_first)
    right_list = append(right_list, parsed_second)
  }
  
  // sort the lists in ascending order
  slices.Sort(left_list)
  slices.Sort(right_list)

  // ensure lists are equal length
  if len(left_list) != len(right_list) {
    panic("lists should be equal length")
  }

  var sum int

  // iterate over lists, pop element from both, append diff to sum
  for i := 0; i < len(left_list); i++ {
    result := float64(right_list[i] - left_list[i])
    sum += int(math.Abs(result))
  }

  // part 1
  fmt.Println("=== Final sum:", sum)

  var score int

  for i := 0; i < len(left_list); i++ {
    occurences := get_occurences(right_list, left_list[i]) 
    score += left_list[i] * occurences
  }

  // part 2
  fmt.Println("=== Similarity Score:", score)
}
