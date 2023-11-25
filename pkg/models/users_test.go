package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	var repos = []RepoLanguage{
		{"TestRepo1", []Language{
			{"Python", 1000}}},
		{"TestRepo2", []Language{
			{"Python", 500},
			{"Go", 100},
			{"Javascript", 500},
		}},
		{"TestRepo3", []Language{
			{"Go", 500},
			{"HTMX", 500},
		}},
	}
	user := createUser("testUser", repos)

	var happyPath = func() {
		expectedLanguages := map[string]int{
			"Python": 1500, "Go": 600, "HTMX": 500, "Javascript": 500,
		}
		actualLanguages := user.getAllLanguages()
		assert.Equal(
			t,
			expectedLanguages,
			actualLanguages,
			fmt.Sprintf("Expected languages %v, received %v", expectedLanguages, actualLanguages),
		)

		expectedLoc := 3100
		actualLoc := user.getTotalLoc()
		assert.Equal(
			t,
			expectedLoc,
			actualLoc,
			fmt.Sprintf("Expected loc %d, received %d", expectedLoc, actualLoc),
		)

		expectedTopLanguages := []Language{
			{"Python", 1500},
			{"Go", 600},
		}
		actualTopLanguages := user.getTopLanguages(2)
		assert.Equal(
			t,
			expectedTopLanguages,
			actualTopLanguages,
			fmt.Sprintf("Expected topLanguages %v, received %v", expectedTopLanguages, actualTopLanguages),
		)
	}

    happyPath()
}
