package main

import (
	"fmt"

	"github.com/michelececcacci/codeforces"
)

func main() {
	c := codeforces.NewClient("", "")
	resp, err := c.Contest.Hacks(566)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}