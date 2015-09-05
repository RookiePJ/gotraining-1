## Arrays - Arrays, Slices and Maps

Arrays are a special data structure in Go that allow us to allocate contiguous blocks of fixed size memory. Arrays have some special features in Go related to how they are declared and viewed as types.

## Notes

* Arrays are fixed length data structures that can't change.
* Arrays of different sizes are considered to be of different types.
* Memory is allocated as a contiguous block.

## Code Review

[Declare, initialize and iterate](example1/example1.go) ([Go Playground](http://play.golang.org/p/2D24t6fbW_))

[Different type arrays](example2/example2.go) ([Go Playground](http://play.golang.org/p/nYgwqqOctt))

[Contiguous memory allocations](example3/example3.go) ([Go Playground](http://play.golang.org/p/L-SmdGfUcP))

## Exercises

### Exercise 1

Declare an array of 5 strings with each element initialized to its zero value. Declare a second array of 5 strings and initialize this array with literal string values. Assign the second array to the first and display the results of the first array.

[Template](exercises/template1/template1.go) ([Go Playground](http://play.golang.org/p/ggjjRPzhAB)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](http://play.golang.org/p/Pa3mrTCcpB))
___
All material is licensed under the [GNU Free Documentation License](https://github.com/gobridge/gotraining/blob/master/LICENSE).