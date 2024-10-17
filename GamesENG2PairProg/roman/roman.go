//Authors
// Shane Conroy , Jack Nulty and Darragh McKernan

package main

//imports for code
import (
	"fmt"
	"log"
)

// main func
func main() {
	var input string                     // input value to be given by user
	fmt.Print("Enter a Roman Numeral: ") // user msg to enter roman vlaue
	_, err := fmt.Scanln(&input)         // scan input
	if err != nil {                      // check error on user input
		log.Fatal(err)
	}
	number := generateNumbersFromString(input)   // send user input to generate number from the roman string
	fmt.Println("The integer value is:", number) // print value to user
}

func generateNumbersFromString(Roman string) int {
	var answer int               // answer var
	var numbers []int            // array of possible roman values
	romanValues := map[rune]int{ // map the values of the roman characters
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, char := range Roman { // loop through each character in the user input and add the numbers to the numbers array
		if value, exists := romanValues[char]; exists {
			numbers = append(numbers, value)
		} // end if
	} // end for
	for i := 0; i < len(numbers); i++ { // for each number in the array of numbers loop through
		if i < len(numbers)-1 && numbers[i] < numbers[i+1] { // if the current is less than the next number then take away the current number from the answer
			answer -= numbers[i]
		} else { // else add the number
			answer += numbers[i]
		} // end else
	}
	return answer
}
