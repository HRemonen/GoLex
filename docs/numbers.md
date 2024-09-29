# Numbers in Lox

The numbers in Lox are all evaluated to 64 bit floating points at runtime. Regardless this, both integers and decimal numbers are supported. Number literal is a series of digits followed by a `.`and one or more trailing digits.

```go
1234
1234.56
```

By design, Lox does not allow a leading or trailing decimal point:

```go
.1234 // Invalid
1234. // Invalid
```

This is to keep the logic simple as of now atleast. This could be something to work on at a later point of time.

## Considerations

Supporting trailing decimal point might cause issues if we decided to supply methods on numbers such as `123.sqrt()`. Also to recognize the leading decimal point, we would have to look at any digit following the `COMMA` character, which would be rather tedious also.