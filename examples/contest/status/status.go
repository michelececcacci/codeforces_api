// nolint
package main

import (
	"fmt"

	"github.com/michelececcacci/codeforces"
)

func main() {
	c, err := codeforces.NewCustomClient()
	resp, err := c.Contest.StatusWithHandle(566, 1, 2, "cheeto1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
