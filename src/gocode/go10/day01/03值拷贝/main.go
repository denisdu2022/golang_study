package main

import (
	"fmt"
	"strings"
)

func main() {

	var s = "mysql ... -u root -p pwd"
	s0 := strings.Contains(s, "root")
	fmt.Println(s0)
	s1 := strings.Contains(s, "root")
	fmt.Println(s1)
	fmt.Println(strings.Contains(s, "pwd"))

}
