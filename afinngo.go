package afinngo

import (
	"regexp"
	"strings"
)

const (
	dictionaryNotFound = "Dictionary could not be loaded successfully."
)

// ScoringStrategy is an interface that holds a function to determine a scoring strategy
type ScoringStrategy interface {
	getScore(tokens []string, currentTokenIndex int, currentTokenScore float64) float64
	getTokens() map[string]float64
	getNegators() map[string]float64
}

// SentimentAnalyzer is loaded with a dictionary and contains the method to provide sentiment analysis
type SentimentAnalyzer struct {
	Strategy ScoringStrategy
	Tokens   map[string]float64
	Negators map[string]float64
}

// SentimentResult contains the results of a sentiment analysis
type SentimentResult struct {
	Score       float64
	Comparative float64
	Tokens      []string
	Words       []string
	Positive    []string
	Negative    []string
}

func tokenize(input string) []string {
	// We only care about hyphens, spaces, and alphanumerics for tokenization
	symbolReplacer := regexp.MustCompile("[^a-zA-Z0-9'\\s-]+")

	return strings.Split(
		strings.ToLower(symbolReplacer.ReplaceAllLiteralString(input, "")),
		" ",
	)
}

// NewDefaultSentimentAnalyzer returns an English AFINN165 sentiment analyzer
func NewDefaultSentimentAnalyzer() *SentimentAnalyzer {
	strategy := NewEnglishStrategy()

	return &SentimentAnalyzer{
		Strategy: strategy,
		Tokens:   strategy.getTokens(),
		Negators: strategy.getNegators(),
	}
}

// NewSentimentAnalyzerFromStrategy loads a dictionary and returns a new instance of a SentimentAnalyzer, given a strategy
func NewSentimentAnalyzerFromStrategy(strategy ScoringStrategy) *SentimentAnalyzer {
	return &SentimentAnalyzer{
		Strategy: strategy,
		Tokens:   strategy.getTokens(),
		Negators: strategy.getNegators(),
	}
}

// Analyze a phrase for its sentiment (positive or negative)
func (a *SentimentAnalyzer) Analyze(phrase string) SentimentResult {
	result := SentimentResult{
		Positive:    []string{},
		Negative:    []string{},
		Words:       []string{},
		Score:       0.0,
		Comparative: 0.0,
		Tokens:      tokenize(phrase),
	}

	for i, token := range result.Tokens {
		// If the word has no AFINN value, skip it
		tokenScore, tokenInDictionary := a.Tokens[token]

		if !tokenInDictionary {
			continue
		}

		tokenScore = a.Strategy.getScore(result.Tokens, i, tokenScore)

		result.Words = append(result.Words, token)

		if tokenScore > 0 {
			result.Positive = append(result.Positive, token)
		}

		if tokenScore < 0 {
			result.Negative = append(result.Negative, token)
		}

		result.Score += tokenScore
	}

	numberOfTokens := len(result.Tokens)

	if numberOfTokens > 0 {
		result.Comparative = result.Score / float64(numberOfTokens)
	}

	return result
}
