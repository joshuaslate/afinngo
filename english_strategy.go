package afinngo

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func buildDictionary(language, dictionaryType string) (map[string]float64, error) {
	dictionary := make(map[string]float64)
	usableLanguage := language

	if usableLanguage == "" {
		usableLanguage = "en"
	}

	dictionaryPath := fmt.Sprintf("./dictionaries/%s/%s.csv", usableLanguage, dictionaryType)
	csvFile, err := os.Open(dictionaryPath)
	defer csvFile.Close()

	if err != nil {
		return dictionary, errors.New(dictionaryNotFound)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			break
		}

		dictionary[line[0]], _ = strconv.ParseFloat(line[1], 64)
	}

	return dictionary, nil
}

// EnglishStrategy is the default scoring strategy (implements ScoringStrategy interface)
type EnglishStrategy struct {
	Tokens   map[string]float64
	Negators map[string]float64
}

// NewEnglishStrategy builds the default sentiment scoring strategy for the English language
func NewEnglishStrategy() *EnglishStrategy {
	languageCode := "en"
	// Build the tokens dictionary
	tokens, tokenErr := buildDictionary(languageCode, "tokens")
	// Build the negators dictionary
	negators, negatorErr := buildDictionary(languageCode, "negators")

	if tokenErr != nil || negatorErr != nil {
		return nil
	}

	strategy := &EnglishStrategy{
		Tokens:   tokens,
		Negators: negators,
	}
	return strategy
}

func (s *EnglishStrategy) getTokens() map[string]float64 {
	return s.Tokens
}

func (s *EnglishStrategy) getNegators() map[string]float64 {
	return s.Negators
}

// getScore(tokens []string, currentTokenIndex int) float64
func (s *EnglishStrategy) getScore(tokens []string, currentTokenIndex int, currentTokenScore float64) float64 {
	tokenScore := currentTokenScore

	if currentTokenIndex > 0 {
		prevToken := tokens[currentTokenIndex-1]
		// If the previous token was a negation word, we will subtract the current word's score rather than adding it
		if _, isNegationWord := s.Negators[prevToken]; isNegationWord {
			tokenScore = -currentTokenScore
		}
	}

	return tokenScore
}
