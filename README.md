# go-calculator
The `go-calculator` package provides a function called `Calculate` with the ability to perform basic math calculations with a string.

## Table of Contents
---
- [go-calculator](#go-calculator)
	- [Table of Contents](#table-of-contents)
	- [Install](#install)
	- [How does it work?](#how-does-it-work)
	- [Calculate](#calculate)
		- [Main File](#main-file)
		- [Output](#output)
	- [Examples](#examples)
		- [Some Basic Calculations](#some-basic-calculations)
		- [Some Complex Calculations](#some-complex-calculations)
	- [License](#license)
---

## Install
Use `go get` to install this package:

```sh
go get github.com/brianlewyn/go-calculator@v1.0.0
```

## How does it work?

Firstly, each character is assigned a new identity, for example, reading the number `π` as a `rune` (equal to `int32`), I take it as `TokenKind` (equal to `uint8`). I do this to minimize memory usage. It should be noted that the numbers do not acquire a new identity, on the contrary, their form is maintained.

Secondly, each time a new identity is assigned it is saved in a 'linked list'. I do this to have a left-to-right reading control.

Next, I check the grammar of the expression.

Finally, I solve in order of hierarchy. But, I only convert the numbers when I have to perform some operation, otherwise I keep them as a string.

## Calculate
The Calculate function receives a string as a parameter and returns a possible result and a possible error.

### Main File
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

### Some Basic Calculations

| Calculations | Results |
| :----------- | :------ |
| 2            | 2       |
| 5+5          | 10      |
| 5%2          | 1       |
| 5^3          | 125     |
| √√625        | 5       |

###  Some Complex Calculations
| Calculations                                 | Results            |
| :------------------------------------------- | :----------------- |
| (1+2+(3+4+(5+6+(7+8)+9+10)+11+12)+13+14)     | 105                |
| (-1-2-(-3-4-(-5-6-(-7-8)-9-10)-11-12)-13-14) | -15                |
| √(2^2 + 3^2 + 4^2 + 5^4) + √(2*π) + 15*6 -1  | 117.08005197971984 |


## License
This package is licensed under `MIT License`. See the LICENSE file for details.