/******************************************************************************
 * Radio Code Calculator API usage example.
 *
 * Version      : v1.1.6
 * Language     : Go
 * Author       : Bartosz Wójcik
 * Homepage     : https://www.pelock.com
 *
 *****************************************************************************/

package main

import (
	"fmt"

	radiocodecalculator "github.com/PELock/Radio-Code-Calculator-Go"
)

func main() {
	model := radiocodecalculator.FordMSeries
	fmt.Println("valid 123456 ->", model.Validate("123456", ""))
	fmt.Println("bad length ->", model.Validate("123", ""))
	fmt.Println("bad pattern ->", model.Validate("ABCDEF", ""))
}
