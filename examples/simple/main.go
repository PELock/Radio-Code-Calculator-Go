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
	"context"
	"fmt"
	"os"

	radiocodecalculator "github.com/PELock/Radio-Code-Calculator-Go"
)

func main() {
	client := radiocodecalculator.New("ABCD-ABCD-ABCD-ABCD")
	errCode, result, err := client.Calc(context.Background(), radiocodecalculator.FordMSeries, "123456", "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch errCode {
	case radiocodecalculator.ErrorSuccess:
		fmt.Println("Radio code is", result["code"])
	case radiocodecalculator.ErrorInvalidLicense:
		fmt.Println("Invalid license key!")
	default:
		fmt.Printf("Error code %d\n", errCode)
	}
}
