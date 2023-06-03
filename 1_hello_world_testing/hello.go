package hello

import "fmt"

func hello() string {
	return "Hello, World!"
}

func main() {
	// It is good to separate your "domain" code from the outside world (side-effects).
	// The fmt.Println is a side effect (printing to stdout) and the string we send in
	// is our domain. 
	// fmt.Println("Hello, World!")

	// so I've created a function called hello() that returns a string
	// and I'm calling that function here
	fmt.Println(hello())
}
