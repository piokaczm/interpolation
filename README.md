GoLang lib for calculating interpolations

!!! matrix should work just fine, other methods WIP !!!

Since my use required it, there are some constraints:
* it returns map in format xi: f(xi) or array of f(x)'s
* it returns only values for x int > 0
* it rounds f(x) and returns them as ints

However it should be relatively easy to add start-end points to struct for iteration over range with negative numbers, or add function for returning value for given argument.


Feel free to contribute

Thanks to gonum team for providing mat64 lib
