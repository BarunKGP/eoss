package models

import (
	"log"
	"time"
)

type User struct {
	Username      string
	CreatedAt     time.Time
	RepoLanguages []RepoLanguage
}

func createUser(uname string, repoLanguages []RepoLanguage) *User {
    if len(uname) == 0 {
        log.Fatal("Username must have atleast 1 character")
    }
    if len(repoLanguages) == 0 {
        log.Fatal("No repository languages passed")
    }

	user := User{Username: uname, CreatedAt: time.Now(), RepoLanguages: repoLanguages}

	return &user
}

func (u User) getAllLanguages() map[string] int {
	langLoc := make(map[string] int)
	for _, item := range u.RepoLanguages {
		for _, lang := range item.Languages {
			if _, ok := langLoc[lang.Name]; !ok {
				langLoc[lang.Name] = 0
			}

			langLoc[lang.Name] += lang.Loc
		}
	}

    return langLoc
}

func (u User) getTotalLoc() int {
    langLoc := u.getAllLanguages()
    var totalLoc int
    for _, loc := range langLoc{
        totalLoc += loc
    }

    return totalLoc
}

func (u User) getTopLanguages(k int) []Language {
	// Parse `RepoLanguages` to find total lines of code and LOC for each language
	sumLoc := u.getTotalLoc()
	langLoc := u.getAllLanguages()

	// Find top `k` languages according to LOC
    var langSlice = MapToLanguages(langLoc)
    SortTopLanguages(langSlice)

	var topLangs []Language
	for i := 0; i < k; i++ {
		log.Printf("Rank %d.  %s -> %f", i+1, langSlice[i].Name, float32(langSlice[i].Loc/sumLoc))
		topLangs = append(topLangs, langSlice[i])
	}

	return topLangs
}
