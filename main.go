package main

import (
	"fmt"
	"github.com/grinny853/puppy"
)
func main() {
	s := puppy.Bark()

	fmt.Println(s)

	s1 := puppy.Barks()

	fmt.Println(s1)

	s2 := puppy.BigBark()

	fmt.Println(s2)
}
