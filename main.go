package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jdkato/prose/v2"
)

func main() {
	
	read := readFile("C:\\Users\\Cristina\\Desktop\\create-read-files\\text.txt")
	summary := resume(read)
	createFile(summary)
	
	fmt.Println("Texto transferido e resumido com sucesso!")
}


type SentenceScore struct{
	sentence string
	score float64
}


func resume(text string) string {
	// analisa o arquivo
	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatalf("analitcs text error %v", err)
	}

	// verificar frequencia de palavras
	wordFreq := make(map[string]int)
	for _, tok := range doc.Tokens() {
		word := strings.ToLower(tok.Text)
		if len(word) > 2 {
			wordFreq[word]++
		}
	}

	var sentenceScores []SentenceScore
	for _, sent := range doc.Sentences() {
		score := 0.0
		words := strings.Fields(strings.ToLower(sent.Text))

		for _, word := range words {
			score += float64(wordFreq[word])
		}
		sentenceScores = append(sentenceScores, SentenceScore{sent.Text, score})
	}

	sort.Slice(sentenceScores, func(i, j int) bool {
		return sentenceScores[i].score > sentenceScores[j].score
	})

	numSentences := 3
	if len(sentenceScores) < 3 {
		numSentences = len(sentenceScores)
	}

	summary := ""
	for i := 0; i < numSentences; i++ {
		summary += sentenceScores[i].sentence + "\n" // Add a newline character after each sentence
	}

	return summary
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func createFile(content string) {
	file, err := os.Create("C:\\Users\\Cristina\\Desktop\\create-read-files\\result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []byte(content)
	nb, err := file.Write(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Foram escritos %d bytes\n", nb)
}


