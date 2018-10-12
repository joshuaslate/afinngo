package afinngo

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	expectedOutput := []string{
		"the",
		"quick",
		"brown",
		"fox",
		"jumps",
		"over",
		"the",
		"lazy",
		"dog",
		"let's",
		"test",
		"an",
		"apostrophe",
	}
	output := tokenize("The quick brown fox jumps over the lazy dog. Let's test an apostrophe.")

	for i, currOutput := range output {
		if currOutput != expectedOutput[i] {
			t.Errorf("Tokenization: Expected to receive: %s, but got: %s", expectedOutput, output)
			return
		}
	}
}

func TestPositiveAnalyze(t *testing.T) {
	analyzer := NewDefaultSentimentAnalyzer()
	output := analyzer.Analyze("There is nothing quite like a warm, sunny day at the beach!")

	if output.Comparative <= 0 {
		t.Error("Expected to receive a comparative greater than 0, but instead received: ", output.Comparative)
	}
}

func TestNegativeAnalyze(t *testing.T) {
	analyzer := NewDefaultSentimentAnalyzer()
	output := analyzer.Analyze("Nobody likes a blood-sucking mosquito. They're truly awful creatures.")

	if output.Comparative >= 0 {
		t.Error("Expected to receive a comparative less than 0, but instead received: ", output.Comparative)
	}
}

func TestNegation(t *testing.T) {
	analyzer := NewDefaultSentimentAnalyzer()
	outputWithoutNegation := analyzer.Analyze("I am happy.")
	outputWithNegation := analyzer.Analyze("I am not happy.")

	if outputWithoutNegation.Comparative <= 0 {
		t.Error("Expected to receive a comparative greater than 0, but instead received: ", outputWithoutNegation.Comparative)
	}

	if outputWithNegation.Comparative >= 0 {
		t.Error("Expected to receive a comparative less than 0, but instead received: ", outputWithNegation.Comparative)
	}
}
