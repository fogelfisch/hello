package main

import "fmt"
import "rsc.io/quote"

// ü

func main() {
	fmt.Println("Hallowelt")
	fmt.Println("Hello(): " + quote.Hello())
	fmt.Println("Go(): "    + quote.Go())
    fmt.Println("Opt(): "   + quote.Opt())
    fmt.Println("Glass(): " + quote.Glass())
}
