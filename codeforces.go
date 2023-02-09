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
	Actions  *actionsService
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
		Actions: &actionsService{c},
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
	randomPrefix := fmt.Sprint(randomInRange(1e5, 1e6))
	text := (randomPrefix + "/" + suffix + "?" + oldParams + "#" + c.apiSecret)
	hash := sha512.Sum512([]byte(text))
	params.Add("apiSig", randomPrefix+fmt.Sprintf("%x", hash))
	base.RawQuery = params.Encode()
	resp, err := c.client.Get(base.String())
	return resp, err
}

type service struct {
	client *httpClientWrapper
}

type (
	blogService    service
	userService    service
	contestService service
	problemService service
	actionsService service
)

func (s *blogService) Comments(id uint) (*[]Comment, error) {
	params := map[string]string{"blogEntryId": fmt.Sprint(id)}
	resp, err := s.client.Get("blogEntry.comments", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[[]Comment]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *blogService) EntryById(id uint) (*BlogEntry, error) {
	params := map[string]string{"blogEntryId": fmt.Sprint(id)}
	resp, err := s.client.Get("blogEntry.view", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[BlogEntry]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *contestService) Hacks(id uint) (*ContestHack, error) {
	params := map[string]string{"contestId": fmt.Sprint(id)}
	resp, err := s.client.Get("contest.hacks", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[ContestHack]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *contestService) List(gym bool) (*[]Contest, error) {
	params := map[string]string{"gym": fmt.Sprint(gym)}
	resp, err := s.client.Get("contest.list", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[[]Contest]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, err
}

func (s *userService) Info(users []string) (*[]User, error) {
	params := map[string]string{"handles": strings.Join(users, ";")}
	resp, err := s.client.Get("user.info", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[[]User]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *userService) Rating(user string) (*[]RatingChange, error) {
	params := map[string]string{"handle": user}
	resp, err := s.client.Get("user.rating", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[[]RatingChange]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

func (s *contestService) Standings(contestId, from, count uint, handles []string, room, unofficial bool) (*ContestStandings, error) {
	params := map[string]string{
		"contestId":  fmt.Sprint(contestId),
		"from":       fmt.Sprint(from),
		"count":      fmt.Sprint(count),
		"handles":    strings.Join(handles, ";"),
		"room":       fmt.Sprint(room),
		"unofficial": fmt.Sprint(unofficial),
	}
	resp, err := s.client.Get("contest.standings", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[ContestStandings]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
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
		"handle":    handle,
	}
	resp, err := s.client.Get("contest.status", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[[]ContestStatus]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

// tags can also be empty (will return every problem in the problemset)
func (s *problemService) Problemset(tags []string) (*Problemset, error) {
	params := map[string]string{"tags": strings.Join(tags, ";")}
	resp, err := s.client.Get("problemset.problems", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[Problemset]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

// Maximum count can be up to 100
func (s *actionsService) RecentActions(count uint) (*[]RecentAction, error) {
	if count > 100 {
		return nil, fmt.Errorf("Count is greater than 100")
	}
	params := map[string]string{"maxCount": fmt.Sprint(count)}
	resp, err := s.client.Get("recentActions", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[[]RecentAction]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}

// Requires authentication
func (s *userService) Friends(onlyOnline bool) (*[]string, error) {
	params := map[string]string{"onlyOnline": fmt.Sprint(onlyOnline)}
	resp, err := s.client.Get("user.friends", params)
	err = handleResponseStatusCode(resp, err)
	if err != nil {
		return nil, err
	}
	rw := ResultWrapper[[]string]{}
	err = unmarshalToResultWrapper(&rw, resp.Body)
	if err != nil {
		return nil, err
	}
	return &rw.Result, nil
}
