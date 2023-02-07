// nolint
package main

import (
	"fmt"

	"github.com/michelececcacci/codeforces"
)

func main() {
	c := codeforces.NewClient("", "")
	resp, err := c.User.Info([]string{"tourist"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
