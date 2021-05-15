package main

import (
	"fmt"
	"github.com/mopore/go-hello-world/morestrings"
	"github.com/google/go-cmp/cmp"
)


func main() {
	fmt.Println( morestrings.ReverseRunes("!oG ,olleH") )
	fmt.Println( cmp.Diff("Hello World", "Hello Go") )
}
