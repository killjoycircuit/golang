//var ->to create variable 
var i :=10

//func ->to declare a function
func

//defer ->to delay the execution of a function until the surrounding function returns ,Deferred calls execute even if a panic occurs (unless the program crashes)
package main

import "fmt"

func main() {
    fmt.Println("Start")

    defer fmt.Println("This runs at the end")

    fmt.Println("Middle")
}

