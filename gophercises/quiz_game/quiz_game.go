package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var score int
	// Use flags package to show cli options
	csvFile := readArguments()

	// Read the csv
	questions := make(map[string]string)
	questions = createQuestionsMap(*csvFile) // get the value from the pointer

	// Print each question
	reader := bufio.NewReader(os.Stdin)
	for question := range questions {
		fmt.Printf("Question: %s= ?\n", question)

		// Wait for user input
		answer, _ := reader.ReadString('\n')

		// Validate the input
		if strings.Compare(strings.Trim(strings.ToLower(answer), "\n"), questions[question]) == 0 {
			score++
		}
	}

	fmt.Printf("\nYOUR SCORE: %d\n", score)
}

func createQuestionsMap(filePath string) map[string]string {
	questions := make(map[string]string)
	// Load the csv file
	raw, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	// Create a new reader
	reader := csv.NewReader(raw)
	for {
		line, err := reader.Read()

		if len(line) >= 1 {
			questions[line[0]] = line[1]
		}

		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

	return questions
}

func readArguments() *string {
	csvFile := flag.String("csv", "questions.csv", "CSV file with the question for the quiz, they must be (question, value) formated.")

	flag.Parse()

	return csvFile
}
