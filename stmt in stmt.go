/*package main

import (
	"fmt"
	"strconv"
)

func main() {

	n, err := strconv.Atoi("42")

	if err == nil {
		fmt.Println("there was no error , n is ", n)
	}
}*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("there was no error , n is ", n)
	}
}





















