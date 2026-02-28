package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	//go:embed assets/nouns.txt
	nouns embed.FS

	//go:embed assets/verbs.txt
	verbs embed.FS

	//go:embed assets/adjectives.txt
	adjectives embed.FS

	//go:embed assets/adverbs.txt
	adverbs embed.FS

	//go:embed assets/conjunctions.txt
	conjunctions embed.FS

	nounsSlice, verbsSlice, adjectivesSlice, adverbsSlice, conjunctionsSlice []string
)

func init() {
	nounsSlice = loadWords(nouns, "assets/nouns.txt")
	verbsSlice = loadWords(verbs, "assets/verbs.txt")
	adjectivesSlice = loadWords(adjectives, "assets/adjectives.txt")
	adverbsSlice = loadWords(adverbs, "assets/adverbs.txt")
	conjunctionsSlice = loadWords(conjunctions, "assets/conjunctions.txt")
}

func main() {
	useID := flag.Bool("i", false, "generate a random character ID instead of a human-readable ID")
	separator := flag.String("s", "-", "word separator for human-readable IDs")
	flag.Parse()

	count := 4
	if len(flag.Args()) > 1 {
		fmt.Fprintln(os.Stderr, "usage: humanid [-i] [count]")
		os.Exit(1)
	}
	if len(flag.Args()) == 1 {
		parsedCount, err := strconv.Atoi(flag.Arg(0))
		if err != nil || parsedCount <= 0 {
			fmt.Fprintln(os.Stderr, "count must be at least 1")
			os.Exit(1)
		}
		count = parsedCount
	}

	if *useID {
		fmt.Println(GenerateID(count))
		return
	}

	fmt.Println(strings.ReplaceAll(GenerateHumanID(count), "-", *separator))
}

func GenerateID(size int) string {
	if size < 1 {
		size = 1
	}
	readableChars := "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
	readableLen := len(readableChars)
	id := make([]byte, size)
	for i := range id {
		id[i] = readableChars[rand.Intn(readableLen)]
	}
	return string(id)
}

func GenerateHumanID(numWords int) string {
	if numWords <= 1 {
		numWords = 1
	}

	for {
		remaining := numWords
		parts := make([]string, 0, (numWords+4)/5)
		for remaining >= 6 {
			parts = append(parts, generateHumanIDUpToSix(6))
			remaining -= 5
		}
		if remaining > 0 {
			parts = append(parts, generateHumanIDUpToSix(remaining))
		}

		id := strings.Join(parts, "-")
		if isValid(id) {
			return id
		}
	}
}

func generateHumanIDUpToSix(numWords int) string {
	patterns := map[int][][]string{
		1: {nounsSlice},
		2: {adjectivesSlice, nounsSlice},
		3: {adverbsSlice, adjectivesSlice, nounsSlice},
		4: {adjectivesSlice, nounsSlice, verbsSlice, adverbsSlice},
		5: {adverbsSlice, adjectivesSlice, nounsSlice, verbsSlice, adverbsSlice},
		6: {adjectivesSlice, nounsSlice, verbsSlice, adverbsSlice, conjunctionsSlice},
	}

	wordPattern := patterns[4]
	if mappedPattern, ok := patterns[numWords]; ok {
		wordPattern = mappedPattern
	}

	id := ""
	for !isValid(id) {
		words := make([]string, 0, len(wordPattern))
		for _, source := range wordPattern {
			words = append(words, pickRandom(source))
		}
		id = strings.Join(words, "-")
	}

	return strings.ReplaceAll(id, " ", "-")
}

func isValid(input string) bool {
	if input == "" {
		return false
	}

	return true
	/*
		substrings := strings.Split(input, "-")
		counts := make(map[string]int)
		for _, word := range substrings {
			counts[word]++
			if counts[word] > 1 {
				return false
			}
		}
		return true

	*/
}

func loadWords(fs embed.FS, filename string) []string {
	data, err := fs.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to load words: %s", err.Error()))
	}

	data = bytes.TrimSpace(data)
	lines := strings.Split(string(data), "\n")
	words := make([]string, 0, len(lines))
	for _, word := range lines {
		if strings.Contains(word, " ") {
			continue
		}
		words = append(words, word)
	}

	return words
}

func pickRandom(slice []string) string {
	return slice[rand.Intn(len(slice))]
}
