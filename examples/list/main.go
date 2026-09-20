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
	errCode, models, err := client.List(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if errCode != radiocodecalculator.ErrorSuccess {
		fmt.Printf("Error code %d\n", errCode)
		os.Exit(1)
	}
	for _, m := range models {
		fmt.Printf("%s serial_len=%d extra_len=%d\n", m.Name, m.SerialMaxLen, m.ExtraMaxLen)
	}
}
