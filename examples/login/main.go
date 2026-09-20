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
	errCode, result, err := client.Login(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if errCode != radiocodecalculator.ErrorSuccess {
		fmt.Printf("Error code %d\n", errCode)
		os.Exit(1)
	}
	license, _ := result["license"].(map[string]any)
	fmt.Println("License activation status -", license["activationStatus"])
	fmt.Println("License owner -", license["userName"])
	fmt.Println("License type -", license["type"])
	fmt.Println("Expiration date -", license["expirationDate"])
}
