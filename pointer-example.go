package main

import "fmt"

func calculate(x, y, a int) int {
  return x + y + a
}

func n (n *int) {
  *n = 9
}

func main() {
  x := 6
  y := 2
  a := 7
  r := calculate(x, y, a)
  fmt.Println(r)
  n(&x)
  fmt.Println(x)
  s := calculate(x, y, a)
  fmt.Println(s)
}