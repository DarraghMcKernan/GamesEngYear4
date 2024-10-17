// Authors
// Shane Conroy , Jack Nulty and Darragh McKernan
// Pair programed in brain waves whoever had an idea took the wheel (Jesus wasnt available)
// Description
/*
	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
	For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
	Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.
	Given a roman numeral, convert it to an integer.

	Example 1:
	Input: s = "III"
	Output: 3
	Explanation: III = 3.
	Example 2:
	Input: s = "LVIII"
	Output: 58
	Explanation: L = 50, V= 5, III = 3.
	Example 3:
	Input: s = "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

	Constraints:
	1 <= s.length <= 15
	s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
	It is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/

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
