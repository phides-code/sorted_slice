/*Write a program which prompts the user to enter integers and stores the integers in a 
sorted slice. The program should be written as a loop. Before entering the loop, the program 
should create an empty integer slice of size (length) 3. During each pass through the loop, 
the program prompts the user to enter an integer to be added to the slice. The program adds 
the integer to the slice, sorts the slice, and prints the contents of the slice in sorted 
order. The slice must grow in size to accommodate any number of integers which the user 
decides to enter. The program should only quit (exiting the loop) when the user enters the 
character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var mySlice = make([]int, 3)
	var userInput string
	var myInt int
	var y int

	for i := 0; ; i++ {

		fmt.Printf("\nPlease enter an integer (X to quit): ")
		fmt.Scan(&userInput)

		if strings.ToLower(userInput) == "x" {
			break
		}

		myInt, _ = strconv.Atoi(userInput)

		if i <= 2 { // we're still in the first 3 ints
			switch { // this switch ensures no negative ints are overwritten
			case mySlice[0] == 0:
				mySlice[0] = myInt

			case mySlice[1] == 0:
				mySlice[1] = myInt

			default:
				mySlice[2] = myInt
			}
		} else { // we're beyond 3 ints and growing the slice
			mySlice = append(mySlice, myInt)
		}

		sort.Ints(mySlice)
		fmt.Printf("The slice now contains:\n")

		for _, y = range mySlice {
			fmt.Printf("%d\n", y)
		}
	}

	fmt.Printf("\nExiting.")
}
