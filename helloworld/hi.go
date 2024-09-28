package main

import "fmt"
import "rsc.io/quote"
import "testlib"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())
    fmt.Println(testlib.Hello("wuforde"))
}
