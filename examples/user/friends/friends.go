package main

import (
	"fmt"
	"os"

	"github.com/michelececcacci/codeforces"
)

func main() {
	key := os.Getenv("CF_API_KEY")
	secret := os.Getenv("CF_API_SECRET")
	c := codeforces.NewClient(key, secret)
	resp, err := c.User.Friends(false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
