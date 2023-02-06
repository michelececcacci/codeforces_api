package codeforces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	cases := []struct {
		user            User
		currentDivision uint
		maxDivision     uint
	}{
		{user: User{Rating: 10, MaxRating: 1000}, currentDivision: 4, maxDivision: 4},
		{user: User{Rating: 1000, MaxRating: 2000}, currentDivision: 4, maxDivision: 2},
		{user: User{Rating: 1200, MaxRating: 1200}, currentDivision: 4, maxDivision: 4},
		{user: User{Rating: 1399, MaxRating: 1400}, currentDivision: 4, maxDivision: 3},
		{user: User{Rating: 1400, MaxRating: 1800}, currentDivision: 3, maxDivision: 2},
		{user: User{Rating: 1500, MaxRating: 1500}, currentDivision: 3, maxDivision: 3},
		{user: User{Rating: 1599, MaxRating: 1599}, currentDivision: 3, maxDivision: 3},
		{user: User{Rating: 1600, MaxRating: 1600}, currentDivision: 2, maxDivision: 2},
		{user: User{Rating: 1800, MaxRating: 3000}, currentDivision: 2, maxDivision: 1},
		{user: User{Rating: 1800, MaxRating: 1899}, currentDivision: 2, maxDivision: 2},
		{user: User{Rating: 1900, MaxRating: 1900}, currentDivision: 2, maxDivision: 2},
		{user: User{Rating: 3200, MaxRating: 3200}, currentDivision: 1, maxDivision: 1},
	}
	for _, tt := range cases {
		assert.Equal(t, tt.currentDivision, tt.user.CurrentDivision())
		assert.Equal(t, tt.maxDivision, tt.user.MaxDivision())
	}
}

func TestIsRated(t *testing.T) {
	rated := User{Rating: 2}
	unrated := User{Rating: 0}
	assert.True(t, rated.IsRated())
	assert.False(t, unrated.IsRated())
}
