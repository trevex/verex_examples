package main

import (
	"fmt"
	"github.com/trevex/verex"
)



func main() {	
	r := verex.New().Find("CaSe").WithAnyCase().Compile()
	
	fmt.Println(r.MatchString("case"))
	fmt.Println(r.MatchString("CASE"))
	fmt.Println(r.MatchString("WUT?"))
}
