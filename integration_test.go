package codeforces

// Currently used to verify the api endpoints are hit.
// Still really early stages

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	c Client
}

func (suite *IntegrationSuite) SetupTest() {
	suite.c = *NewClient("", "")
}

func (suite *IntegrationSuite) TestInfo() {
	resp, err := suite.c.User.Info([]string{"tourist"})
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
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

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}
