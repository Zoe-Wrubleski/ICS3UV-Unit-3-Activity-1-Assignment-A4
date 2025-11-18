/*
 * Author Zoe Wrubleski
 * Version 1.0.0
 * Date 2025-11-18
 * Fileoverview This program shows the average of 3 numbers
 */

package main

import "fmt"

func main() {
	// set veriables
	var number1 float32 = 56.9
	var number2 float32 = 89.7
	var number3 float32 = 90.2

	// calculate average
	var answer float32 = (number1 + number2 + number3) / 3

	// display answer
	fmt.Println("The average of 56.9, 89.7, and 90.2 is", answer)

	fmt.Println("\nDone")
}
