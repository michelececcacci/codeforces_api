package codeforces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentDivision(t *testing.T) {
	cases := []struct {
		user            User
		currentDivision uint
	}{
		{user: User{Rating: 10}, currentDivision: 4},
		{user: User{Rating: 1000}, currentDivision: 4},
		{user: User{Rating: 1200}, currentDivision: 4},
		{user: User{Rating: 1399}, currentDivision: 4},
		{user: User{Rating: 1400}, currentDivision: 3},
		{user: User{Rating: 1500}, currentDivision: 3},
		{user: User{Rating: 1599}, currentDivision: 3},
		{user: User{Rating: 1600}, currentDivision: 2},
		{user: User{Rating: 1800}, currentDivision: 2},
		{user: User{Rating: 1800}, currentDivision: 2},
		{user: User{Rating: 1900}, currentDivision: 2},
		{user: User{Rating: 3200}, currentDivision: 1},
	}
	for _, tt := range cases {
		assert.Equal(t, tt.currentDivision, tt.user.CurrentDivision())
	}
}

func TestMaxDivision(t *testing.T) {
	cases := []struct {
		user User
		maxDivision uint
	} {
		{user: User{MaxRating: 1000}, maxDivision: 4},
		{user: User{MaxRating: 2000}, maxDivision: 2},
		{user: User{MaxRating: 1200}, maxDivision: 4},
		{user: User{MaxRating: 1400}, maxDivision: 3},
		{user: User{MaxRating: 1800}, maxDivision: 2},
		{user: User{MaxRating: 1500}, maxDivision: 3},
		{user: User{MaxRating: 1599}, maxDivision: 3},
		{user: User{MaxRating: 1600}, maxDivision: 2},
		{user: User{MaxRating: 3000}, maxDivision: 1},
		{user: User{MaxRating: 1899}, maxDivision: 2},
		{user: User{MaxRating: 1900}, maxDivision: 2},
		{user: User{MaxRating: 3200}, maxDivision: 1},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.maxDivision, tt.user.MaxDivision())
	}
}

func TestIsRated(t *testing.T) {
	rated := User{Rating: 2}
	unrated := User{Rating: 0}
	assert.True(t, rated.IsRated())
	assert.False(t, unrated.IsRated())
}
