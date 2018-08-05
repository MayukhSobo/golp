package vectorize

import (
	"fmt"
	"reflect"
	"strings"
)

// FindUniqueWordsFromList -- Find the common words in corpus in List
func FindUniqueWordsFromList(ignoreCaps bool, corpus ...string) {

	// Create a list of all the list of corpus and tokensize the word
	corpusList := make([][]string, len(corpus))
	for i := 0; i < len(corpus); i++ {
		var parts []string
		if ignoreCaps {
			parts = strings.Split(strings.ToLower(corpus[i]), " ")
		} else {
			parts = strings.Split(corpus[i], " ")
		}
		corpusList[i] = parts
	}

	// Create a map to hold the counts of all the words
	uniqueWords := map[string]bool{}
	for _, doc := range corpusList {
		for _, word := range doc {
			uniqueWords[word] = true
		}
	}

	fmt.Println((reflect.ValueOf(uniqueWords).MapKeys()))
}
