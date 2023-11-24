package models

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoLanguages(t *testing.T) {
	var langs = []Language{
		{"Python", 100},
		{"Java", 100},
		{"Go", 200},
	}

	var happyPath = func() {
		repo := RepoLanguage{"Test Repo", langs}
		expectedLangs := []string{"Python", "Java", "Go"}
		actualLangs := repo.getLanguages()

		expectedLen := 3
		actualLen := repo.Len()

		expectedTotalLoc := 400
		actualTotalLoc := repo.getTotalLoc()

		assert.Equal(t, expectedLen, actualLen, fmt.Sprintf("Returned len %d, Expected %d", actualLen, expectedLen))
		assert.Equal(t, expectedLangs, actualLangs, fmt.Sprintf("Returned languages %v, Expected %v", actualLangs, expectedLangs))
		assert.Equal(t, expectedTotalLoc, actualTotalLoc, fmt.Sprintf("Returned languages %d, Expected %d", actualTotalLoc, expectedTotalLoc))
	}

	var zeroLangPath = func() {
		repo := RepoLanguage{"ZeroLanguageRepo", []Language{}}

		var expectedLangs = []string(nil)
		actualLangs := repo.getLanguages()

		expectedLen := 0
		actualLen := repo.Len()

		expectedTotalLoc := 0
		actualTotalLoc := repo.getTotalLoc()

		assert.Equal(t, expectedLen, actualLen, fmt.Sprintf("Returned len %d, Expected %d", actualLen, expectedLen))
		assert.Equal(t, expectedLangs, actualLangs, fmt.Sprintf("Returned languages %v, Expected %v", actualLangs, expectedLangs))
		assert.Equal(t, expectedTotalLoc, actualTotalLoc, fmt.Sprintf("Returned languages %d, Expected %d", actualTotalLoc, expectedTotalLoc))

	}

	log.Println("Running happyPath")
	happyPath()

	log.Println("Running zeroLanguagesPath")
	zeroLangPath()
}

func TestLanguageConversion(t *testing.T) {
    langSlice := []Language{
        {"Python", 100},
		{"Java", 100},
		{"Go", 200},
    }

    langMap := map[string] int{"Python": 100, "Java": 100, "Go": 200}

    actualLangMap := LanguagesToMap(langSlice)
    actualLangSlice := MapToLanguages((langMap))

    assert.Equal(
        t, 
        langSlice, 
        actualLangSlice, 
        fmt.Sprintf("Returned list %v, Expected %v", actualLangSlice, langSlice),
    )
    assert.Equal(
        t, 
        langMap, 
        actualLangMap, 
        fmt.Sprintf("Returned map %v, Expected %v", actualLangMap, langMap),
    )
}