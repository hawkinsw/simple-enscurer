### Simple Enscurer

A tool (suitable for use with `//go:generate` to use a shift cipher
to _enscure_ string literals in your code.

The tool will automatically generate a getter function (whose name you
choose) that will return the string you specify on the command line without
having that string literal in the binary.


