// Word Frequency takes a file as input and output the frequency of
// each word that appears in the file.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type FreqMap map[string]uint64

type FreqPair struct {
	text  string
	count uint64
}

type FreqPairs []*FreqPair

// Implement sort.Interface to make FreqPairs sortable
func (pairs FreqPairs) Len() int {
	return len(pairs)
}

func (pairs FreqPairs) Swap(i, j int) {
	pairs[i], pairs[j] = pairs[j], pairs[i]
}

// Different types of sorting wrappers
type ByText struct{ FreqPairs }

func (pairs ByText) Less(i, j int) bool {
	return pairs.FreqPairs[i].text < pairs.FreqPairs[j].text
}

type ByCount struct{ FreqPairs }

func (pairs ByCount) Less(i, j int) bool {
	return pairs.FreqPairs[i].count > pairs.FreqPairs[j].count
}

// Run opens the file located at path and printout the frequency output.
// Run returns an error if there's any failure.
// The frequency output is:
// 	<word>: <freq count>
// The output can be sorted either by frequency or by alphabetically order
func Run(path string) error {
	fmt.Println("file:", path)

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	freqs := make(FreqMap)
	for scanner.Scan() {
		text := scanner.Text()
		_, exists := freqs[text]
		if exists {
			freqs[text]++
		} else {
			freqs[text] = 1
		}
	}

	fmt.Println(CreateFreqPairs(freqs))

	return nil
}

func CreateFreqPairs(freqs FreqMap) FreqPairs {
	pairs := make(FreqPairs, len(freqs))
	i := 0
	for text, count := range freqs {
		pairs[i] = &FreqPair{text, count}
		i++
	}

	sort.Sort(ByText{pairs})
	return pairs
}

func (pair *FreqPair) String() string {
	return fmt.Sprintf("%v: %v\n", pair.text, pair.count)
}
