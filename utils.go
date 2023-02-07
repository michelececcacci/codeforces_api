package codeforces

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func unmarshalToResultWrapper[T any](rw *ResultWrapper[T], reader io.Reader) error {
	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &rw)
	if err != nil {
		return err
	}
	return nil
}

func handleResponseStatusCode(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprint(resp.StatusCode))
	}
	return nil
}

// non inclusive integer in range
func randomInRange(min, max int) int {
	return min+ rand.Intn(max-min)
}
