package main

import (
	"fmt"
)

/*
romanToInt converts a Roman numeral string to an integer.

Solution Explanation:
The function iterates through each character of the Roman numeral string,
adding its respective value to the total sum. For characters that can
represent subtractive combinations (e.g., IV, IX), the function subtracts
twice the value of the previous character when a smaller numeral precedes
a larger numeral (e.g., when 'I' precedes 'V' or 'X') to correct the total.

The logic is implemented using a switch statement that checks each character
and adjusts the total sum based on Roman numeral rules.

Time Complexity:
The time complexity of this algorithm is O(n), where n is the length of the 
Roman numeral string. This is because the function processes each character 
exactly once in a single pass through the string.

Space Complexity:
The space complexity of this algorithm is O(1), as it uses a fixed amount of
extra space regardless of the input size. Only a few integer variables are 
used to store the intermediate results and the previous character.
*/

func romanToInt(s string) int {
	// Initialize variables: 'prev' to track the previous character,
	// and 'sum' to accumulate the total integer value.
	var prev rune
	var sum int
	
	// Iterate over each character in the string 's'
	for _, v := range s {
		// Use a switch statement to handle each Roman numeral character.
		switch v {
		case 'I': // Case for Roman numeral 'I'
			sum += 1
		case 'V': // Case for Roman numeral 'V'
			sum += 5
			/*
			If the previous character was 'I', subtract 2 to correct for over-counting.
			This handles the case where 'I' precedes 'V' (IV = 4).
			*/
			if prev == 'I' {
				sum -= 2
			}
		case 'X': // Case for Roman numeral 'X'
			sum += 10
			/*
			If the previous character was 'I', subtract 2 to correct for over-counting.
			This handles the case where 'I' precedes 'X' (IX = 9).
			*/
			if prev == 'I' {
				sum -= 2
			}
		case 'L': // Case for Roman numeral 'L'
			sum += 50
			/*
			If the previous character was 'X', subtract 20 to correct for over-counting.
			This handles the case where 'X' precedes 'L' (XL = 40).
			*/
			if prev == 'X' {
				sum -= 20
			}
		case 'C': // Case for Roman numeral 'C'
			sum += 100
			/*
			If the previous character was 'X', subtract 20 to correct for over-counting.
			This handles the case where 'X' precedes 'C' (XC = 90).
			*/
			if prev == 'X' {
				sum -= 20
			}
		case 'D': // Case for Roman numeral 'D'
			sum += 500
			/*
			If the previous character was 'C', subtract 200 to correct for over-counting.
			This handles the case where 'C' precedes 'D' (CD = 400).
			*/
			if prev == 'C' {
				sum -= 200
			}
		case 'M': // Case for Roman numeral 'M'
			sum += 1000
			/*
			If the previous character was 'C', subtract 200 to correct for over-counting.
			This handles the case where 'C' precedes 'M' (CM = 900).
			*/
			if prev == 'C' {
				sum -= 200
			}
		}
		// Update 'prev' to the current character for the next iteration.
		prev = v
	}
	// Return the calculated sum which is the integer representation of the Roman numeral.
	return sum
}

func main() {
	// Example usage of the romanToInt function.
	roman := "MCMXCIV"
	integer := romanToInt(roman)
	fmt.Printf("The integer value of %s is %d\n", roman, integer)
}
