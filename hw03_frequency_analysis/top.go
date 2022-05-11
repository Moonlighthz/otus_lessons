package hw03frequencyanalysis

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

var re = regexp.MustCompile(`[[:punct:]]`)

func Top10(text string) []string {
	slice := strings.Fields(text)
	var result []Word

	for _, word := range slice {
		word, e := clearString(word)

		if e != nil {
			continue
		}

		if i, e := findResultStruct(result, word); e == nil {
			result[i].count++
			continue
		}

		result = append(result, Word{word, 1})
	}

	return getTopWord(sortSlice(result))
}

func clearString(s string) (string, error) {
	s = strings.ToLower(s)
	ar := strings.Split(s, "-")

	for i := range ar {
		ar[i] = re.ReplaceAllString(ar[i], "")
	}

	s = strings.Join(ar, "-")

	if s == "-" {
		return "", errors.New("uncorrect string")
	}

	return strings.Join(ar, "-"), nil
}

func sortSlice(s []Word) []Word {
	sort.Slice(s, func(i, j int) bool {
		if s[i].count > s[j].count {
			return true
		}

		if s[i].count < s[j].count {
			return false
		}

		return s[i].word < s[j].word
	})

	return s
}

func getTopWord(result []Word) []string {
	var list []string
	for i, word := range result {
		if i > 9 {
			break
		}
		list = append(list, word.word)
	}

	return list
}

func findResultStruct(result []Word, word string) (int, error) {
	for i := range result {
		if result[i].word == word {
			return i, nil
		}
	}

	return -1, errors.New("element not founded")
}
