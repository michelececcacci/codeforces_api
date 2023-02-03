package codeforces

import (
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntryByIdValid(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte(`
		{
			"status": "OK",
			"result": {
				"originalLocale": "ru",
				"allowViewHistory": false,
				"creationTimeSeconds": 1267562173,
				"rating": 14,
				"authorHandle": "MikeMirzayanov",
				"modificationTimeSeconds": 1267651613,
				"id": 123,
				"title": "Codeforces Maintenance",
				"locale": "en",
				"tags": [
				"codeforces",
				"maintenance"
				]
			}
		}`))
		assert.Nil(t, err)
	}))
	defer ts.Close()
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	bs := blogService{c}
	resp, err := bs.EntryById(123)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "ru", resp.OriginalLocale)
	assert.False(t, resp.AllowViewHistory)
	assert.Equal(t, 1267562173, resp.CreationTimeSeconds)
	assert.Equal(t, 14, resp.Rating)
	assert.Equal(t, "MikeMirzayanov", resp.AuthorHandle)
	assert.Equal(t, 1267651613, resp.ModificationTimeSeconds)
	assert.Equal(t, 123, resp.ID)
	assert.Equal(t, "Codeforces Maintenance", resp.Title)
	assert.Equal(t, "en", resp.Locale)
	assert.Equal(t, []string{"codeforces", "maintenance"}, resp.Tags)
}

func TestEntryByIdInvalid(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	}))
	defer ts.Close()
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	bs := blogService{c}
	resp, err := bs.EntryById(0)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

func TestHacks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte(`{
				"status": "OK",
				"result": [
					{
					"id": 160426,
					"creationTimeSeconds": 1438274514,
					"hacker": {
						"contestId": 566,
						"members": [
						{
							"handle": "Sehnsucht"
						}
						],
						"participantType": "CONTESTANT",
						"ghost": false,
						"room": 29,
						"startTimeSeconds": 1438273200
					},
					"defender": {
						"contestId": 566,
						"members": [
						{
							"handle": "osama"
						}
						],
						"participantType": "CONTESTANT",
						"ghost": false,
						"room": 29,
						"startTimeSeconds": 1438273200
					},
					"verdict": "INVALID_INPUT",
					"problem": {
						"contestId": 566,
						"index": "F",
						"name": "Clique in the Divisibility Graph",
						"type": "PROGRAMMING",
						"points": 500.0,
						"rating": 1500,
						"tags": [
						"dp",
						"math",
						"number theory"
						]
					},
					"judgeProtocol": {
						"protocol": "Validator 'validate.exe' returns exit code 3 [FAIL Integer parameter [name=a[1]] equals to 2, violates the range [3, 1000000] (stdin)]",
						"manual": "false",
						"verdict": "Invalid input"
					}
					},
					{
					"id": 160427,
					"creationTimeSeconds": 1438274878,
					"hacker": {
						"contestId": 566,
						"members": [
						{
							"handle": "Misha100896"
						}
						],
						"participantType": "CONTESTANT",
						"ghost": false,
						"room": 5,
						"startTimeSeconds": 1438273200
					},
					"defender": {
						"contestId": 566,
						"members": [
						{
							"handle": "fruwajacybyk"
						}
						],
						"participantType": "CONTESTANT",
						"ghost": false,
						"room": 5,
						"startTimeSeconds": 1438273200
					},
					"verdict": "HACK_UNSUCCESSFUL",
					"problem": {
						"contestId": 566,
						"index": "F",
						"name": "Clique in the Divisibility Graph",
						"type": "PROGRAMMING",
						"points": 500.0,
						"rating": 1500,
						"tags": [
						"dp",
						"math",
						"number theory"
						]
					},
					"judgeProtocol": {
						"protocol": "Solution verdict:\nOK\n\nChecker:\nok 1 number(s): \"2\"\r\n\n\nInput:\n2\r\n1 2\r\n\n\nOutput:\n2\r\n\n\nAnswer:\n2\n\n\nTime:\n46\n\nMemory:\n9080832\n",
						"manual": "false",
						"verdict": "Unsuccessful hacking attempt"
					}
					}
				]
		}`))
		assert.Nil(t, err)
	}))
	defer ts.Close()
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	cs := contestService{c}
	resp, err := cs.Hacks(566)
	firstHacker := Hacker{
		Members:          []Member{{Handle: "Sehnsucht"}},
		ContestID:        566,
		ParticipantType:  "CONTESTANT",
		Ghost:            false,
		Room:             29,
		StartTimeSeconds: 1438273200}

	assert.Len(t, *(resp), 2)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 160426, (*resp)[0].ID)
	assert.Equal(t, firstHacker, (*resp)[0].Hacker)
}

// TODO check that raw query is actually the right one
func TestInfoSingleUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_ , err := w.Write([]byte(`{
			"status":"OK",
			"result":[
				{
					"lastName":"Korotkevich",
					"country":"Belarus",
					"lastOnlineTimeSeconds":1675068776,
					"city":"Gomel",
					"rating":3803,
					"friendOfCount":52616,
					"titlePhoto":"https://userpic.codeforces.org/422/title/50a270ed4a722867.jpg",
					"handle":"tourist",
					"avatar":"https://userpic.codeforces.org/422/avatar/2b5dbe87f0d859a2.jpg",
					"firstName":"Gennady",
					"contribution":147,
					"organization":"ITMO University",
					"rank":"legendary grandmaster",
					"maxRating":3979,
					"registrationTimeSeconds":1265987288,
					"maxRank":"legendary grandmaster"}]
				}`))
		assert.Nil(t, err)
	}))
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	defer ts.Close()
	us := userService{c}
	resp, err := us.Info([]string{"tourist"})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, (*resp), 1)
	assert.Equal(t, "Korotkevich", (*resp)[0].LastName)
	assert.Equal(t, 3803, (*resp)[0].Rating)
	assert.Equal(t, 1675068776, (*resp)[0].LastOnlineTimeSeconds)
	assert.Equal(t, "Gomel", (*resp)[0].City)
}

func TestInfoMultipleUsers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte(`{
			"status": "OK",
			"result": [
				{
				"lastName": "Korotkevich",
				"country": "Belarus",
				"lastOnlineTimeSeconds": 1675068776,
				"city": "Gomel",
				"rating": 3803,
				"friendOfCount": 52618,
				"titlePhoto": "https://userpic.codeforces.org/422/title/50a270ed4a722867.jpg",
				"handle": "tourist",
				"avatar": "https://userpic.codeforces.org/422/avatar/2b5dbe87f0d859a2.jpg",
				"firstName": "Gennady",
				"contribution": 147,
				"organization": "ITMO University",
				"rank": "legendary grandmaster",
				"maxRating": 3979,
				"registrationTimeSeconds": 1265987288,
				"maxRank": "legendary grandmaster"
				},
				{
				"lastName": "Qi",
				"country": "United States",
				"lastOnlineTimeSeconds": 1675044771,
				"city": "Princeton",
				"rating": 3783,
				"friendOfCount": 11172,
				"titlePhoto": "https://userpic.codeforces.org/312472/title/7cf0a442d4071e87.jpg",
				"handle": "Benq",
				"avatar": "https://userpic.codeforces.org/312472/avatar/5716ac69aea8159a.jpg",
				"firstName": "Benjamin",
				"contribution": 50,
				"organization": "MIT",
				"rank": "legendary grandmaster",
				"maxRating": 3813,
				"registrationTimeSeconds": 1435099979,
				"maxRank": "legendary grandmaster"
				}
			]
		}`))
		assert.Nil(t, err)
	}))
	defer ts.Close()
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	us := userService{c}
	resp, err := us.Info([]string{"tourist", "benq"})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, (*resp), 2)
	tourist := (*resp)[0]
	benq := (*resp)[1]
	assert.Equal(t, "Korotkevich", tourist.LastName)
	assert.Equal(t, "Belarus", tourist.Country)
	assert.Equal(t, "Gennady", tourist.FirstName)
	assert.Equal(t, "Qi", benq.LastName)
	assert.Equal(t, "United States", benq.Country)
}

func TestRating(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte(`{
			"status": "OK",
			"result": [
				{
				"contestId": 2,
				"contestName": "Codeforces Beta Round #2",
				"handle": "tourist",
				"rank": 14,
				"ratingUpdateTimeSeconds": 1267124400,
				"oldRating": 0,
				"newRating": 1602
				},
				{
				"contestId": 8,
				"contestName": "Codeforces Beta Round #8",
				"handle": "tourist",
				"rank": 5,
				"ratingUpdateTimeSeconds": 1270748700,
				"oldRating": 1602,
				"newRating": 1764
				}
			]
		}`))
		assert.Nil(t, err)
	}))
	defer ts.Close()
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	us := userService{c}
	firstResult := RatingChange{
		ContestID:               2,
		ContestName:             "Codeforces Beta Round #2",
		Handle:                  "tourist",
		Rank:                    14,
		RatingUpdateTimeSeconds: 1267124400,
		OldRating:               0,
		NewRating:               1602,
	}
	resp, err := us.Rating("tourist")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, (*resp), 2)
	assert.Equal(t, (*resp)[0], firstResult)
}

func TestStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err :=  w.Write([]byte(`{
			"status": "OK",
			"result": [
				{
				"id": 12291750,
				"contestId": 566,
				"creationTimeSeconds": 1438347312,
				"relativeTimeSeconds": 2147483647,
				"problem": {
					"contestId": 566,
					"index": "A",
					"name": "Matching Names",
					"type": "PROGRAMMING",
					"points": 1750,
					"rating": 2300,
					"tags": [
					"dfs and similar",
					"strings",
					"trees"
					]
				},
				"author": {
					"contestId": 566,
					"members": [
					{
						"handle": "tourist"
					}
					],
					"participantType": "PRACTICE",
					"ghost": false,
					"startTimeSeconds": 1438273200
				},
				"programmingLanguage": "GNU C++11",
				"verdict": "OK",
				"testset": "TESTS",
				"passedTestCount": 38,
				"timeConsumedMillis": 171,
				"memoryConsumedBytes": 29388800
				}
			]
		}`))
		assert.Nil(t, err)
	}))
	defer ts.Close()
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	cs := contestService{c}
	resp, err := cs.Status(566, 1, 2, "tourist")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, *resp, 1)
	assert.Equal(t, (*resp)[0].ID, 12291750)
	assert.Equal(t, (*resp)[0].ContestID, 566)
	assert.Equal(t, (*resp)[0].Problem.Index, "A")
}

// Even though API seems to return tags in sorted order, we just ignore that
// and sort both input and output for an easier comparison
func TestProblemSet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte(`{
			"status": "OK",
			"result": {
				"problems": [
				{
					"contestId": 1535,
					"index": "B",
					"name": "Array Reodering",
					"type": "PROGRAMMING",
					"rating": 900,
					"tags": [
					"brute force",
					"greedy",
					"math",
					"number theory",
					"sortings"
					]
				}
				],
				"problemStatistics": [
				{
					"contestId": 1535,
					"index": "B",
					"solvedCount": 25512
				}
				]
			}
		}`))
		assert.Nil(t, err)
	}))
	c := newDefaultClientWrapper(ts.URL+"/", "", "")
	ps := problemService{c}
	tags := []string{
		"greedy",
		"math",
		"number theory",
		"brute force",
		"sortings"}
	resp, err := ps.Problemset(tags)
	problem := Problem{
		ContestID: 1535,
		Index:     "B",
		Name:      "Array Reodering",
		Type:      "PROGRAMMING",
		Rating:    900,
		Tags:      tags,
	}
	// we don't really care about order of tags, so we can just sort both
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	sort.Strings(tags)
	sort.Strings(problem.Tags)
	assert.Len(t, resp.Problems, 1)
	assert.Len(t, resp.ProblemStatistics, 1)
	assert.Equal(t, problem, resp.Problems[0])
}
