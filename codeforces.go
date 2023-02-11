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
	return serializeResponse[[]Comment](resp, err)
}

func (s *blogService) EntryById(id uint) (*BlogEntry, error) {
	params := map[string]string{"blogEntryId": fmt.Sprint(id)}
	resp, err := s.client.Get("blogEntry.view", params)
	return serializeResponse[BlogEntry](resp, err)
}

func (s *contestService) Hacks(id uint) (*ContestHack, error) {
	params := map[string]string{"contestId": fmt.Sprint(id)}
	resp, err := s.client.Get("contest.hacks", params)
	return serializeResponse[ContestHack](resp, err)
}

func (s *contestService) RatingChange(id uint) (*[]RatingChange, error) {
	params := map[string]string{"contestId": fmt.Sprint(id)}
	resp, err := s.client.Get("contest.ratingChanges", params)
	return serializeResponse[[]RatingChange](resp, err)
}

func (s *contestService) List(gym bool) (*[]Contest, error) {
	params := map[string]string{"gym": fmt.Sprint(gym)}
	resp, err := s.client.Get("contest.list", params)
	return serializeResponse[[]Contest](resp, err)
}

func (s *userService) Info(users []string) (*[]User, error) {
	params := map[string]string{"handles": strings.Join(users, ";")}
	resp, err := s.client.Get("user.info", params)
	err = handleResponseStatusCode(resp, err)
	return serializeResponse[[]User](resp, err)
}

func (s *userService) Rating(user string) (*[]RatingChange, error) {
	params := map[string]string{"handle": user}
	resp, err := s.client.Get("user.rating", params)
	return serializeResponse[[]RatingChange](resp, err)
}

func (s *contestService) Standings(contestId, from, count uint, handles []string, unofficial bool) (*ContestStandings, error) {
	params := map[string]string{
		"contestId":      fmt.Sprint(contestId),
		"from":           fmt.Sprint(from),
		"count":          fmt.Sprint(count),
		"handles":        strings.Join(handles, ";"),
		"showUnofficial": fmt.Sprint(unofficial),
	}
	resp, err := s.client.Get("contest.standings", params)
	return serializeResponse[ContestStandings](resp, err)
}

func statusDefaultParams(contestId, from, count uint) *map[string]string {
	params := map[string]string{
		"contestId": fmt.Sprint(contestId),
		"from":      fmt.Sprint(from),
		"count":     fmt.Sprint(count),
	}
	return &params
}

func (s *contestService) StatusWithHandle(contestId, from, count uint, handle string) (*[]ContestStatus, error) {
	params := statusDefaultParams(contestId, from, count)
	(*params)["handle"] = handle
	resp, err := s.client.Get("contest.status", *params)
	return serializeResponse[[]ContestStatus](resp, err)
}

func (s *contestService) Status(contestId, from, count uint) (*[]ContestStatus, error) {
	resp, err := s.client.Get("contest.status", *statusDefaultParams(contestId, from, count))
	return serializeResponse[[]ContestStatus](resp, err)
}

// tags can also be empty (will return every problem in the problemset)
func (s *problemService) Problemset(tags []string) (*Problemset, error) {
	params := map[string]string{"tags": strings.Join(tags, ";")}
	resp, err := s.client.Get("problemset.problems", params)
	return serializeResponse[Problemset](resp, err)
}

// Maximum count can be up to 100
func (s *actionsService) RecentActions(count uint) (*[]RecentAction, error) {
	if count > 100 {
		return nil, fmt.Errorf("Count is greater than 100")
	}
	params := map[string]string{"maxCount": fmt.Sprint(count)}
	resp, err := s.client.Get("recentActions", params)
	return serializeResponse[[]RecentAction](resp, err)
}

// Requires authentication
func (s *userService) Friends(onlyOnline bool) (*[]string, error) {
	params := map[string]string{"onlyOnline": fmt.Sprint(onlyOnline)}
	resp, err := s.client.Get("user.friends", params)
	return serializeResponse[[]string](resp, err)
}
