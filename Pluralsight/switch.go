package main

import "fmt"

func main() {

	i := 5
	fmt.Println( "I is", i )

	switch i {
	case 1:
		fmt.Println( "This is a" )
		fmt.Println( "one" )
	case 2:
		fmt.Println( "This is a" )
		fmt.Println( "two" )
	default:
		fmt.Println( "This is default" )
	}
}
