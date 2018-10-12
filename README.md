# AfinnGo
An AFINN-165-powered sentiment analysis tool written in Go. This is essentially a port of the npm sentiment module, but in Go for a small side-project I am working on.

## Use
To access the default English AFINN-165 analyzer, you can use the following:

```go
import "github.com/joshuaslate/afinngo"

func isSentencePositive(sentence string) bool {
  analyzer := afinngo.NewDefaultSentimentAnalyzer()
  output := analyzer.Analyze("There is nothing quite like a warm, sunny day at the beach!")

  return output.Comparative > 0
}
```

The output of the Analyzer func is:

```go
type SentimentResult struct {
	Score       float64
	Comparative float64
	Tokens      []string
	Words       []string
	Positive    []string
	Negative    []string
}
```

## Customization
You can also use your own analyzation strategy (dictionaries for different languages, handling scoring differently, etc.) by satisfying the `ScoringStrategy` interface.

```go
type ScoringStrategy interface {
	getScore(tokens []string, currentTokenIndex int, currentTokenScore float64) float64
	getTokens() map[string]float64
	getNegators() map[string]float64
}

// Where customStrategy satisfies the ScoringStrategy interface
myAnalyzer := afinngo.NewSentimentAnalyzerFromStrategy(customStrategy)
```

## Contributions
Contributions are welcome. Please feel empowered to fork this repository, make changes, and open a pull request to send your changes back upstream.
