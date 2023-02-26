# Codeforces 

[![Unit tests](https://github.com/michelececcacci/codeforces/actions/workflows/test.yml/badge.svg)](https://github.com/michelececcacci/codeforces/actions/workflows/test.yml) 
[![Go Reference](https://pkg.go.dev/badge/github.com/michelececcacci/codeforces.svg)](https://pkg.go.dev/github.com/michelececcacci/codeforces)
[![codecov](https://codecov.io/github/michelececcacci/codeforces/branch/main/graph/badge.svg?token=E6JT1TXE9D)](https://codecov.io/github/michelececcacci/codeforces)

Implements all the methods mentioned in the [codeforces api](https://codeforces.com/apiHelp).
Creating a client is really simple, all you have to do is:
```go
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
```
You can also leave the key and secret parameters empty, but you  wont be able to access
methods that require authentication such as ` c.Client.Friends() `. 
For examples refer to the [examples] folder

[examples]: /examples
