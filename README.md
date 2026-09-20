# Radio Code Calculator Online & SDK for Go

**[Radio Code Calculator](https://www.pelock.com/products/radio-code-calculator)** is an online service along with [Web API & SDK](https://www.pelock.com/products/radio-code-calculator/sdk) for generating car radio unlock codes for popular vehicle brands.

Following a breakdown or a disconnection of the car battery, most of the vehicle radio & navigation units will ask for an unlocking code. It's standard anti-theft protection.

Our car Radio Code Calculator allows you to generate **100% valid** radio codes to unlock car radios & navigation without the need to use the expensive service of authorized dealers.

![Radio Code Calculator](https://www.pelock.com/img/en/products/radio-code-calculator/car-radio-code-calculator-online-web-api-sdk.jpg)

The service is available through a simple online interface and `Web API`, with multiple `SDK` development libraries for popular programming languages.

Thanks to our solution, you can create, for instance, mobile or web applications that allow for easy generation of radio codes.

## Supported car models and radios

Our service is being continuously developed and new algorithms are gradually added for new car models and their radios.

If a new algorithm is added, you will get automatic and free access to it as part of your current license.

Individual calculators are available on our site as a paid service for end customers. You can verify them and see lists of supported radios on the relevant subpages:

* [Renault & Dacia](https://www.pelock.com/products/renault-and-dacia-car-radio-code-calculator-generator)
* [Toyota ERC](https://www.pelock.com/products/toyota-erc-calculator-radio-unlock-code-generator)
* [Jeep Cherokee](https://www.pelock.com/products/jeep-cherokee-radio-unlock-code-calculator-generator)
* [Ford M Serial](https://www.pelock.com/products/ford-radio-code-m-serial-calculator-generator)
* [Ford V Serial](https://www.pelock.com/products/ford-radio-code-v-serial-calculator-generator)
* [Ford TravelPilot EX, FX & NX](https://www.pelock.com/products/ford-travelpilot-ex-fx-nx-radio-code-generator-calculator)
* [Chrysler Panasonic TM9](https://www.pelock.com/products/chrysler-panasonic-tm9-car-radio-code-calculator-generator)
* [Chrysler Dodge Ram VP2](https://www.pelock.com/products/chrysler-dodge-ram-uconnect-harman-kardon-radio-code)  
* [Fiat Stilo & Bravo Visteon](https://www.pelock.com/products/fiat-stilo-bravo-visteon-radio-code-calculator-generator)
* [Fiat DAIICHI MOPAR](https://www.pelock.com/products/fiat-daiichi-radio-code-calculator-generator)
* [Fiat Continental 250 & 500 VP1/VP2](https://www.pelock.com/products/fiat-250-500-vp1-vp2-radio-code-calculator-generator)  
* [Nissan Glove Box Immobiliser PIN](https://www.pelock.com/products/nissan-glove-box-pin-code-calculator)
* [Eclipse ESN Unlock Code Calculator](https://www.pelock.com/products/eclipse-esn-unlock-code-calculator)
* [Jaguar Alpine](https://www.pelock.com/products/jaguar-alpine-car-radio-unlock-code-calculator)

## Use of radio code calculator

Where and who can use the radio code generation service and make money from code generation?

### ![Android](https://www.pelock.com/img/en/icons/android-32.png) App developers 
The main audience for our software is clearly developers and programmers, either of mobile or desktop applications.

### ![Shoping cart](https://www.pelock.com/img/en/icons/cart-32.png) Online stores 
If you run an online e-commerce store, you can sell radio codes through it using our software solutions.

### ![Car](https://www.pelock.com/img/en/icons/car-32.png) Auto repair shops
We also encourage car repair shops whose customers often use car radio unlocking services.

### ![Person](https://www.pelock.com/img/en/icons/user-32.png) Private individuals
Private individuals will also profit from our solution by generating codes and selling them on car forums or auction sites such as eBay, Craigslist.

### No limits!

You can generate codes **without limitation** with your purchased one year license.

Set your own price for generating a single code and start earning by using **tried and tested** algorithms from a programming language you know.

If you are not a programmer - don't worry. Just use our [online calculator](https://www.pelock.com/products/radio-code-calculator/online).

## Installation

The preferred way to install the Web API SDK is via Go modules.

```bash
go get github.com/PELock/Radio-Code-Calculator-Go
```

The module path is `github.com/PELock/Radio-Code-Calculator-Go`.

## Packages for other programming languages

The installation packages have been uploaded to repositories for several popular programming languages and their source codes have been published on GitHub:

| Repository   | Language | Installation | Package | GitHub |
| ------------ | ---------| ------------ | ------- | ------ |
| ![Packagist repository for PHP and Composer](https://www.pelock.com/img/logos/repo-packagist-composer.png) | PHP | Add the following line to `require` section of your `composer.json` file `"pelock/radio-code-calculator": "*"` | [Packagist](https://packagist.org/packages/pelock/radio-code-calculator) | [Sources](https://github.com/PELock/Radio-Code-Calculator-PHP)
| ![PyPI repository for Python](https://www.pelock.com/img/logos/repo-pypi.png) | Python | Run `pip install radio-code-calculator` | [PyPi](https://pypi.org/project/radio-code-calculator/) | [Sources](https://github.com/PELock/Radio-Code-Calculator-Python)
| ![NPM repository for JavaScript and TypeScript](https://www.pelock.com/img/logos/repo-npm.png) | JavaScript, TypeScript | Run `npm i radio-code-calculator` or add the following to `dependencies` section of your `package.json` file `"dependencies": { "radio-code-calculator": "latest" },` | [NPM](https://www.npmjs.com/package/radio-code-calculator) | [Sources](https://github.com/PELock/Radio-Code-Calculator-JavaScript)

## Usage examples

### Example — `list/main.go`

```go
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
```

### Example — `login/main.go`

```go
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
```

### Example — `offline/main.go`

```go
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
```

### Example — `simple/main.go`

```go
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
```

See the `examples/` directory in this repository for complete samples.

## Got questions?

If you are interested in the Radio Code Calculator Web API or have any questions regarding radio code generator SDK packages, technical or legal issues, or if something is not clear, [please contact me](https://www.pelock.com/contact). I'll be happy to answer all of your questions.

Bartosz Wójcik

* Visit my site at — https://www.pelock.com
* Twitter — https://twitter.com/PELock
* GitHub — https://github.com/PELock