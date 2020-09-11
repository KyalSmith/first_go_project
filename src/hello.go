package hello

import "fmt"
import "rsc.io/quote"

func Hello() string {
    fmt.Println("Hi There :)")
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	return quote.Hello()
}
