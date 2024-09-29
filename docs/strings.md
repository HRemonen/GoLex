# Strings in Go Lox

Strings are a type of Literals in Go Lox which are pretty much only long lexemes which always begin with a special character `"`. Strings are consumed by iterating the characters until the ending `"` is found. Running out of input before that happens is also handled gracefully.

By default and by choice of design, Go Lox allows multiline strings using the same syntax.

As in any other language, strings must be closed with double quotes aswell.
```go
// This will cause runtime error, and is invalid syntax
"Hello world!
```

## Considerations

It might be better to prohibit multiline strings and add functionality for that separately.