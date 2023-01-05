package codeforces

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func UnmarshalToResultWrapper[T any](rw *ResultWrapper[T], reader io.Reader) error {
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

func HandleResponseStatusCode(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprint(resp.StatusCode))
	}
	return nil
}
