// nolint
package main

import (
	"fmt"

	"github.com/michelececcacci/codeforces"
)

func main() {
	c, err := codeforces.NewCustomClient()
	resp, err := c.Blog.EntryById(79)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
