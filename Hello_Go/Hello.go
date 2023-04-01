// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

// Variables

var (
	FirstName string = "Yassine"
	LastName  string = "Touisse"
	Number    int    = 0753732554
)

func main() {
	var k string
	var j int = 40
	i := 42
	fmt.Println(i)
	fmt.Printf("%v , %T \n", j, i)
	k = strconv.Itoa(i)
	fmt.Printf("%v , %T \n", k, k)
	fmt.Printf("%s", FirstName)
}

// Primitives
