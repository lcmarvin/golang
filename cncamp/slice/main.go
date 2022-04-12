package main

import "fmt"

func main() {
	s1 := []string{"p", "o", "e", "m"}
	s2 := s1[2:]
	fmt.Println(s2)
	s2[1] = "t"
	fmt.Println(s1)
	fmt.Println(s2)
}
