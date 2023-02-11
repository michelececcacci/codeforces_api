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

func serializeResponse[T any](resp *http.Response, err error) (*T, error) {
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[T]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

// non inclusive integer in range
func randomInRange(min, max int) int {
	return min + rand.Intn(max-min)
}
