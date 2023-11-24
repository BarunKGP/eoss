package models

import "sort"

type Language struct {
	Name string
	Loc  int
}

func SortTopLanguages(langs []Language) {
	sort.SliceStable(langs, func(i, j int) bool {
        return langs[i].Loc > langs[j].Loc
    })
}

func MapToLanguages(langMap map[string] int) []Language {
    var languages []Language
    for name, loc := range langMap {
        languages = append(languages, Language{name, loc})
    }
    return languages
}

func LanguagesToMap(languages []Language) map[string] int {
    var langMap = make(map[string] int, len(languages))
    for _, l := range languages {
        langMap[l.Name] = l.Loc
    }

    return langMap
}


type RepoLanguage struct {
	Name      string
	Languages []Language
}

func (rl RepoLanguage) Len() int {
	return len(rl.Languages)
}

func (rl RepoLanguage) getLanguages() []string {
	// var languages = make([]string, len(rl.Languages))
	var languages []string
	for _, lang := range rl.Languages {
		languages = append(languages, lang.Name)
	}
	return languages
}

func (rl RepoLanguage) getTotalLoc() int {
	var totalLoc int
	for _, lang := range rl.Languages {
		totalLoc += lang.Loc
	}
	return totalLoc
}
