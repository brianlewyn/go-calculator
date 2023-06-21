# Calculator Package

[![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/brianlewyn/go-calculator)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/brianlewyn/go-calculator/blob/main/LICENSE)

The `go-calculator` package provides a powerful `Calculate` function for performing basic math calculations with a string input.

## Table of Contents

---
- [Installation](#installation)
- [How does it work?](#how-does-it-work)
- [Usage](#usage)
	- [Example](#example)
	- [Output](#output)
- [Examples](#examples)
	- [Basic Calculations](#basic-calculations)
	- [Complex Calculations](#complex-calculations)
- [License](#license)
---

## Installation

To install the latest version of the `go-calculator` package, use the following command:

```sh
go get -u github.com/brianlewyn/go-calculator
```

## How does it work?

The calculator function assigns a unique identity to each character in the input string. Numbers are maintained in their original form, while other characters are treated as tokens to minimize memory usage. These identities are then stored in a linked list, allowing for left-to-right reading control.

The function validates the grammar of the expression and solves the calculations based on the hierarchy of operations. Numbers are converted when necessary, while retaining their original form as strings when not involved in operations.

## Usage

The Calculate function takes a string expression as input and returns the calculated result and any potential errors.

### Example

```go
package main

import (
   "fmt"
   "log"

   "github.com/brianlewyn/go-calculator"
)

func main() {
    expression := "(0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2"
    res64, err := basic.Calculate(expression)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("[ %s ] = %.2f", expression, res64)
}
```

### Output

```sh
[ (0.5 + 4.5 - 1) * 10 * √(6-2) / 4^2 ] = 5.00
```

## Examples

### Basic Calculations

| Calculations | Results |
| :----------- | :------ |
| 2            | 2       |
| 5+5          | 10      |
| 5%2          | 1       |
| 5^3          | 125     |
| √√625        | 5       |

###  Complex Calculations

| Calculations                                 | Results            |
| :------------------------------------------- | :----------------- |
| (1+2+(3+4+(5+6+(7+8)+9+10)+11+12)+13+14)     | 105                |
| (-1-2-(-3-4-(-5-6-(-7-8)-9-10)-11-12)-13-14) | -15                |
| √(2^2 + 3^2 + 4^2 + 5^4) + √(2*π) + 15*6 -1  | 117.08005197971984 |


## License

The Calculator Package is open-source and released under the [MIT License](https://github.com/brianlewyn/go-linked-list/blob/v3/LICENSE). Feel free to use, modify, and distribute this package according to the terms of the license.

Thank you for choosing the Calculator Package. I hope it proves to be a valuable tool in your projects!