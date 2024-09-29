# The Go Lox Language

DISCLAIMER
This is based on the Lox language of "Crafting Interpreters" by Robert Nystrom. For full documentation of the language I would suggest checking the [GitHub repository](https://github.com/munificent/craftinginterpreters) or the [dedicated site](https://craftinginterpreters.com/).

----

Go Lox is dynamically typed. Variables can store values of any type, and a single variable can enven store values of different types at different times.

Operations on values of the wrong type will lead to an error at runtime.

## Data Types

Go Lox supports the following data types:

- *Booleans*. Two boolean values are supported, like in any other language:
  ```go
  true; // Not false
  false; // Not true
  ```
- *Numbers*. Floating point numbers are supported, which means integers and decimal numbers:
  ```go
  123456; // Integer value
  1234.56; // Decimal value
  ```
  
  You can read more information about [numbers](./numbers.md) in Go Lox.
- *Strings*. Strings are text wrapped inside double quotes.
  ```go
  "Hello world!";
  ""; // Empty string
  ```
- *Null*. Null is used for "no value" values. Original JLox features `nil` as the nullish value, but since the intepreter here is writtern in Go, which already has nil, we are going to use `null`.
