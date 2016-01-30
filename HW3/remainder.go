package main

import "fmt"

func main() {
	var largernum int
	var smallernum int
        var remainder int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&smallernum)
	fmt.Print("Please enter a smaller number: ")
	fmt.Scan(&largernum)

        remainder = (largernum % smallernum)

	fmt.Println(remainder)
	
}