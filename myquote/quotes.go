package myquote

import (
	"fmt"
	"rsc.io/quote"
)

func Writeallquotes() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Hello())
}
