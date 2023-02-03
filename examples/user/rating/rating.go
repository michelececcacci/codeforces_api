// nolint
package main

import (
	"fmt"

	"github.com/michelececcacci/codeforces"
)

func main() {
	c := codeforces.NewClient("", "")
	resp, err := c.User.Rating("cheeto1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
