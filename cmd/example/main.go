package main 

import (
  "fmt"
)

func main() {
  fmt.Println("variadic function")
  fmt.Println("=================")

  sum := Sum(0, 1)

  fmt.Printf("sum of (0, 1) is %v", sum)

  sum = Sum()

  fmt.Printf("sum of () is %v", sum)





  
}

func Sum(numbers ...int) (sum int) {
  for _, number := range numbers {
    sum += number
  }
  
  return sum
}
