package shield

import (	
	"log"
	"math"
)

const defaultProb float64 = 1e-11

type shield struct {
	tokenizer 	    	Tokenizer
	store     	    	Store
	useMinCounterWord	bool
	minCounterWord  	int64	
}

// New - new Shield
func New(t Tokenizer, s Store) Shield {
	return &shield{
		tokenizer: t,
		store:     s,
		useMinCounterWord: false,				
	}
}

func (sh *shield) Destroy() {
	err := sh.store.Close()
	if err != nil {
		log.Println(err)
	}
}

func (sh *shield) Learn(class, text string) (err error) {
	return sh.increment(class, text, 1)
}

func (sh *shield) BulkLearn(sets []Set) (err error) {
	return sh.bulkIncrement(sets, 1)
}

func (sh *shield) Forget(class, text string) (err error) {
	return sh.increment(class, text, -1)
}

func (sh *shield) increment(class, text string, sign int64) (err error) {
	if len(class) == 0 {
		panic("invalid class")
	}
	if len(text) == 0 {
		panic("invalid text")
	}
	return sh.bulkIncrement([]Set{Set{Class: class, Text: text}}, sign)
}

func (sh *shield) bulkIncrement(sets []Set, sign int64) (err error) {
	if len(sets) == 0 {
		panic("invalid data set")
	}
	m := make(map[string]map[string]int64)
	for _, set := range sets {
		tokens := sh.tokenizer.Tokenize(set.Text)
		for k := range tokens {
			tokens[k] *= sign
		}
		if w, ok := m[set.Class]; ok {
			for word, count := range tokens {
				w[word] += count
			}
		} else {
			m[set.Class] = tokens
		}
	}
	for class, words := range m {
		// Sitnan patch: Do not consider words if count is less than sh.minCounterWord
		if sh.useMinCounterWord {
			for word, d := range words {
				if d < sh.minCounterWord {
					delete(m[class], word)
				}
			}
		}
		if err = sh.store.AddClass(class); err != nil {
			log.Println(err)
			return
		}
	}
	//log.Println("Total word with freq sent to Redis is: ", len(m))
	return sh.store.IncrementClassWordCounts(m)
}

func getKeys(m map[string]int64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Score (shield) - learn
func (sh *shield) Score(text string) (scores map[string]float64, err error) {

	// Tokenize text
	wordFreqs := sh.tokenizer.Tokenize(text)
	if len(wordFreqs) < 2 {
		return
	}
	words := getKeys(wordFreqs)

	// Get total class word counts
	totals, err := sh.store.TotalClassWordCounts()
	if err != nil {
		return
	}
	classes := getKeys(totals)

	// Get word frequencies for each class
	classFreqs := make(map[string]map[string]int64)
	for _, class := range classes {
		freqs, err2 := sh.store.ClassWordCounts(class, words)
		if err2 != nil {
			err = err2
			return
		}
		classFreqs[class] = freqs
	}

	// Calculate log scores for each class
	logScores := make(map[string]float64, len(classes))

	hasData := false
	for _, class := range classes {
		freqs := classFreqs[class]
		total := totals[class]

		// Because this classifier is not biased, we don't use prior probabilities
		score := float64(0)

		for _, word := range words {
			// Compute the probability that this word belongs to that class
			wordProb := float64(freqs[word]) / float64(total)
			if wordProb == 0 {
				wordProb = defaultProb
			} else {
				hasData = true
			}
			score += math.Log(wordProb)

		}
		logScores[class] = score

	}

	scores = make(map[string]float64, len(classes))
	if !hasData {
		scores["unknown"] = 1.0
		return
	}

	// Normalize the scores	
	scores = make(map[string]float64, len(classes))	
	var max = -math.MaxFloat64
	for _, score := range logScores {
		if score > max {
			max = score
		}
	}
	for class, score := range logScores {
		scores[class] = max/score
	}	
	return
}

func (sh *shield) Classify(text string) (class string, err error) {
	scores, err := sh.Score(text)	
	if err != nil {
		log.Println(err)
		return
	}
	var score float64
	for k, v := range scores {
		if v > score {
			class, score = k, v
		}
	}
	return
}

func (sh *shield) Reset() error {
	return sh.store.Reset()
}
