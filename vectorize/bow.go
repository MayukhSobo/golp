package vectorize

import (
	"fmt"
)

// FindUniquesInCorpus -- Find the common words in corpus
func FindUniquesInCorpus(corpus ...string) {
	for _, docs := range corpus {
		fmt.Println(docs)
	}
}
