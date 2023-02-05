// Utility package that allows access to the codeforces API.
// Some functions require authentication. For more informations on how to
// get your API key and secret, please refer here: https://codeforces.com/apiHelp
package codeforces

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Version              = "v.0.0.1"
	defaultbaseURLString = "https://codeforces.com/api/"
)

// holds a shared httpclient (could change) and the services
// responsible for communicating with the various parts of the api
type Client struct {
	Blog     *blogService
	User     *userService
	Contest  *contestService
	Problems *problemService
}

func NewClient(apiKey, apiSecret string) *Client {
	c := newDefaultClientWrapper(defaultbaseURLString, apiKey, apiSecret)
	return NewCustomClient(apiKey, apiSecret, c)
}

func NewCustomClient(apiKey, apiSecret string, c *httpClientWrapper) *Client {
	return &Client{
		Blog:    &blogService{c},
		User:    &userService{c},
		Contest: &contestService{c},
	}
}

type httpClientWrapper struct {
	client        *http.Client
	baseUrlString string
	apiKey        string
	apiSecret     string
}

func newDefaultClientWrapper(baseUrlString, apiKey, apiSecret string) *httpClientWrapper {
	return &httpClientWrapper{
		baseUrlString: baseUrlString,
		apiKey:        apiKey,
		client:        http.DefaultClient,
		apiSecret:     apiSecret,
	}
}

func (c *httpClientWrapper) Get(suffix string, userParams map[string]string) (*http.Response, error) {
	base, err := url.Parse(c.baseUrlString + suffix)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	for k, v := range userParams {
		params.Add(k, v)
	}
	params.Add("apiKey", c.apiKey)
	t := fmt.Sprint(time.Now().UTC().UnixMilli() / 1000)
	params.Add("time", t)
	oldParams := params.Encode()
	text := ("123456/" + suffix + "?" + oldParams + "#" + c.apiSecret)
	hash := sha512.Sum512([]byte(text))
	params.Add("apiSig", `123456`+fmt.Sprintf("%x", hash)) // TODO CONVERT TO RANDOM
	base.RawQuery = params.Encode()
	resp, err := c.client.Get(base.String())
	return resp, err
}

type service struct {
	client *httpClientWrapper
}

type blogService service
type userService service
type contestService service
type problemService service

func (s *blogService) Comments(id uint) (*[]Comment, error) {
	params := map[string]string{"blogEntryId": fmt.Sprint(id)}
	resp, err := s.client.Get("blogEntry.comments", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[[]Comment]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *blogService) EntryById(id uint) (*BlogEntry, error) {
	params := map[string]string{"blogEntryId": fmt.Sprint(id)}
	resp, err := s.client.Get("blogEntry.view", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[BlogEntry]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *contestService) Hacks(id uint) (*ContestHack, error) {
	params := map[string]string{"contestId": fmt.Sprint(id)}
	resp, err := s.client.Get("contest.hacks", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[ContestHack]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *contestService) List(gym bool) (*[]Contest, error) {
	params := map[string]string{"gym": fmt.Sprint(gym)}
	resp, err := s.client.Get("contest.list", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[[]Contest]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, err
}

func (s *userService) Info(users []string) (*[]User, error) {
	params := map[string]string{"handles": strings.Join(users, ";")}
	resp, err := s.client.Get("info", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[[]User]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *userService) Rating(user string) (*[]RatingChange, error) {
	params := map[string]string{"handle": user}
	resp, err := s.client.Get("user.rating", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[[]RatingChange]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

// TODO TEST
func (s *contestService) Standings(contestId, from, count uint, handles []string, room, unofficial bool) (*ContestStandings, error) {
	params := map[string]string{
		"contestId":  fmt.Sprint(contestId),
		"from":       fmt.Sprint(from),
		"count":      fmt.Sprint(count),
		"handles":    strings.Join(handles, ";"),
		"room":       fmt.Sprint(room),
		"unofficial": fmt.Sprint(unofficial)}
	resp, err := s.client.Get("contest.standings", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[ContestStandings]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *contestService) Status(contestId, from, count uint, handle string) (*[]ContestStatus, error) {
	params := map[string]string{
		"contestId": fmt.Sprint(contestId),
		"from":      fmt.Sprint(from),
		"count":     fmt.Sprint(count),
		"handle":    handle}
	resp, err := s.client.Get("contest.status", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[[]ContestStatus]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

// tags can also be empty (will return every problem in the problemset)
func (s *problemService) Problemset(tags []string) (*Problemset, error) {
	params := map[string]string{"tags": strings.Join(tags, ";")}
	resp, err := s.client.Get("problemset.problems", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[Problemset]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

// Requires authentication
func (s *userService) Friends(onlyOnline bool) (*[]string, error) {
	params := map[string]string{"onlyOnline": fmt.Sprint(onlyOnline)}
	resp, err := s.client.Get("user.friends", params)
	err = HandleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	var rw = ResultWrapper[[]string]{}
	err = UnmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}
