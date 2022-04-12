package main

import "fmt"

func main() {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	for pos, v := range arr {
		if v == "stupid" {
			arr[pos] = "smart"
		}
		if v == "weak" {
			arr[pos] = "strong"
		}
	}
	fmt.Print(arr)
}
