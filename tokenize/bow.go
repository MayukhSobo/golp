package tokenize

import (
	"strings"
)

// FindUniqueWordsFromList -- Find the common words in corpus in List
func FindUniqueWordsFromList(ignoreCaps bool, corpus [][]string) []string {

	// Create a list of all the list of corpus and tokensize the word
	for index, doc := range corpus {
		if ignoreCaps {
			corpus[index] = strings.Split(strings.ToLower(doc[0]), " ")
		} else {
			corpus[index] = strings.Split(doc[0], " ")
		}
	}
	// Create a map to hold the counts of all the words
	uniqueWords := map[string]bool{}
	for _, doc := range corpus {
		for _, word := range doc {
			uniqueWords[word] = true
		}
	}

	// Get the keys as a token of all the unique words
	var token []string
	for k := range uniqueWords {
		token = append(token, k)
	}
	return token
}
