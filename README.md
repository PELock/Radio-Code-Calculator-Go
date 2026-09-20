# Radio Code Calculator — Go Web API SDK

Go client for the [Radio Code Calculator](https://www.pelock.com/products/radio-code-calculator) Web API. Generate unlocking codes for supported car radios.

Endpoint: `https://www.pelock.com/api/radio-code-calculator/v1`

## Installation

```bash
go get github.com/PELock/Radio-Code-Calculator-Go
```

## Usage

```go
package main

import (
	"context"
	"fmt"

	radiocodecalculator "github.com/PELock/Radio-Code-Calculator-Go"
)

func main() {
	client := radiocodecalculator.New("YOUR-WEB-API-KEY")
	errCode, result, err := client.Calc(context.Background(), radiocodecalculator.FordMSeries, "123456", "")
	if err != nil {
		panic(err)
	}
	if errCode == radiocodecalculator.ErrorSuccess {
		fmt.Println(result["code"])
	}
}
```

Commands: `login`, `calc`, `info`, `list`. Built-in models support offline `RadioModel.Validate`. See `examples/`.

## License

Apache-2.0. Copyright Bartosz Wójcik / PELock.
