package codeforces

// Currently used to verify the api endpoints are hit.
// Still really early stages

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	c Client
}

func (suite *IntegrationSuite) SetupTest() {
	apiKey := os.Getenv("CF_API_KEY")
	apiSecret := os.Getenv("CF_API_SECRET")
	suite.c = *NewClient(apiKey, apiSecret)
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (suite *IntegrationSuite) TestInfo() {
	handle := "tourist"
	resp, err := suite.c.User.Info([]string{handle})
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
	first := (*resp)[0]
	assert.Equal(suite.T(), handle, first.Handle)
}

func (suite *IntegrationSuite) TestComments() {
	resp, err := suite.c.Blog.Comments(79)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
}

func (suite *IntegrationSuite) TestHacks() {
	resp, err := suite.c.Contest.Hacks(566)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
}

func (suite *IntegrationSuite) TestEntryByID() {
	resp, err := suite.c.Blog.EntryById(79)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
	assert.Equal(suite.T(), 79, resp.ID)
}

func (suite *IntegrationSuite) TestFriends() {
	suite.showEmptyVariablesWarning()
	resp, err := suite.c.User.Friends(false)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
}

func (suite *IntegrationSuite) showEmptyVariablesWarning() {
	apiKey := os.Getenv("CF_API_KEY")
	apiSecret := os.Getenv("CF_API_SECRET")
	emptyVariableMessage := "%s is empty, TestFriends will probably fail"
	if apiKey == "" {
		suite.T().Logf(emptyVariableMessage, "apiKey")
	}
	if apiSecret == "" {
		suite.T().Logf(emptyVariableMessage, "apiSecret")
	}
}

func (suite *IntegrationSuite) TestStatus() {
	resp, err := suite.c.Contest.StatusWithHandle(566, 1, 5, "tourist")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
}

func (suite *IntegrationSuite) TestStandings() {
	resp, err := suite.c.Contest.Standings(566, 1, 5, []string{}, false)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
}
