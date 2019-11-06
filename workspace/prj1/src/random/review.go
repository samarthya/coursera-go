package main

import "fmt"

/*
func main() {
  x := []int {4, 8, 5}
  y := -1
  for _, elt := range x {
    if elt > y {
      y = elt
    }
  }
  fmt.Print(y)
}
*/
/*
 func main() {
  x := [...]int {4, 8, 5}
  y := x[0:2]
  z := x[1:3]
  y[0] = 1
  z[1] = 3
  fmt.Print(x)
}
*/
/*
func main() {
  x := [...]int {1, 2, 3, 4, 5}
  y := x[0:2]
  z := x[1:4]
  fmt.Print(len(y), cap(y), len(z), cap(z))
}
*/
/*
func main() {
  x := map[string]int {
    "ian": 1, "harris": 2}
  for i, j := range x {
    if i == "harris" {
      fmt.Print(i, j)
    }
  }
}
*/
/*
type P struct {
    x string
y int
}
func main() {
  b := P{"x", -1}
  a := [...]P{P{"a", 10},
        P{"b", 2},
        P{"c", 3}}
  for _, z := range a {
    if z.y > b.y {
      b = z
    }
  }
  fmt.Println(b.x)
}
*/
/*
func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}*/

func main() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}
