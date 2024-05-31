package main

import "fmt"

// Variable declaration with 'var' can be done outside the func
var a string = "This is the longest variable declaration!"
var b = "This is little better"

// Multiple variable declarations
var (
    d = "Hip x"
    e = 2
    f = "Hooray!"
)

var g = "v is used to print value"
var h = "T is used to print type of variable"

func main() {
    c := "This is the shortest, but can only be declared inside functions"
    fmt.Println("Hello, World!")
    fmt.Println("a->" + a + " b->" + b + " c->" + c)
    fmt.Println(d, e, f)
    fmt.Printf("%%%v", g)
    fmt.Printf(" & %%%v.", h)
    fmt.Printf(" Like, variables g & h are %T variables\n", h)
}
