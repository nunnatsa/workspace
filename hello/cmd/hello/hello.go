package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nunnatsa/workspace/hello"
)

func main() {
	var str string
	if len(os.Args) > 1 {
		str = strings.Join(os.Args[1:], " ")
	} else {
		all, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		str = string(all)
	}
	fmt.Println(hello.Reverse(str))
}
