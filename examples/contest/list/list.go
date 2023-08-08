package main

import (
	"fmt"

	"github.com/michelececcacci/codeforces"
)

func main() {
	c, err := codeforces.NewCustomClient()
	resp, err := c.Contest.List(false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
