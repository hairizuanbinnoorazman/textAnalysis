package analyse

import (
	"regexp"
	"sort"
	"strings"
)

func TopNwords(text string, limit int) []Word {
	// Preprocessing the text data
	text = strings.ToLower(text)

	// Split the text data
	stringSlice := splitText(text)

	// Compute word count
	wordCount := topNWordCount(stringSlice, limit)
	return wordCount
}

func splitText(text string) []string {
	r := regexp.MustCompile(`(\w+['-]{1}\w+)|(\w+)`)
	matches := r.FindAllString(text, -1)
	return matches
}

type Word struct {
	Text      string `json:"text"`
	Frequency int    `json:"frequency"`
}

type WordList []Word

func (w WordList) Len() int {
	return len(w)
}

func (w WordList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WordList) Less(i, j int) bool {
	if w[i].Frequency < w[j].Frequency {
		return true
	}
	if w[i].Frequency == w[j].Frequency {
		return w[i].Text > w[j].Text
	}

	return false
}

func PushSortLimit(wordList WordList, word Word, limit int) WordList {
	// Check existence of word in wordList
	wordReplaced := false
	for idx, value := range wordList {
		if value.Text == word.Text {
			wordReplaced = true
			wordList[idx] = word
			break
		}
	}
	if !wordReplaced {
		wordList = append(wordList, word)
	}

	sort.Sort(sort.Reverse(wordList))
	if len(wordList) == limit+1 {
		wordList = wordList[0:limit]
	}
	return wordList
}

func topNWordCount(textSlice []string, limit int) []Word {
	wordMap := make(map[string]int)
	wordSlice := []Word{}
	wordList := WordList(wordSlice)

	for _, value := range textSlice {
		if value == "" {
			continue
		}
		wordMap[value] += 1
		wordList = PushSortLimit(wordList, Word{value, wordMap[value]}, limit)
	}

	return wordList
}
