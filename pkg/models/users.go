package models

import (
	"eoss/pkg/handlers"
	"log"
	"sort"
	"time"
)

type User struct {
	Username      string
	CreatedAt     time.Time
	RepoLanguages []handlers.LanguageType
	TopLanguages  []string
}

func createUser(uname string, repoLanguages []handlers.LanguageType) *User {
	user := User{Username: uname, CreatedAt: time.Now(), RepoLanguages: repoLanguages}
	user.TopLanguages = user.getTopLanguages(5)

	return &user
}

func (u User) getTopLanguages(k int) []string {
	// Parse `RepoLanguages` to find total lines of code and LOC for each language
	var sumLoc int
	langLoc := make(map[string]int)
	for _, item := range u.RepoLanguages {
		for lang, loc := range item.Language {
			if _, ok := langLoc[lang]; !ok {
				langLoc[lang] = 0
			}

			langLoc[lang] += loc
			sumLoc += loc
		}
	}

	// Find top `k` languages according to LOC
	type langSliceType struct {
		Lang    string
		LangLoc int
	}

	var langSlice = make([]langSliceType, len(langLoc))
	for k, v := range langLoc {
		langSlice = append(langSlice, langSliceType{k, v})
	}

	sort.SliceStable(langSlice, func(i, j int) bool {
		return langSlice[i].LangLoc > langSlice[j].LangLoc
	})

	var topLangs []string
	for i := 0; i < k; i++ {
		log.Printf("Rank %d.  %s -> %f", i+1, langSlice[i].Lang, float32(langSlice[i].LangLoc/sumLoc))
		topLangs = append(topLangs, langSlice[i].Lang)
	}

	return topLangs
}
