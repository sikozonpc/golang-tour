package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Problem : Problem struct
type Problem struct {
	Question string
	Answer   string
}

func main() {
	var score int
	// Use flags package to show cli options
	csvFile, timeLimit := readArguments()

	// Read the csv
	lines := readCsvFile(*csvFile) // get the value from the pointer
	problems := parseLines(lines)

	fmt.Printf("All settled for a %v seconds Quiz with %d problems...\n", *timeLimit, len(problems))
	// User input to start the quiz
	var ready string
	fmt.Scanf("Are you ready ? ", &ready)

	if ready == "OK" {
		timer := time.NewTimer(*timeLimit * time.Second)

		// Print each question
		reader := bufio.NewReader(os.Stdin)
		for _, problem := range problems {
			fmt.Printf("Question: %s= ?\n", problem.Question)

			// Wait for user input
			answer, _ := reader.ReadString('\n')

			// Validate the input
			if strings.Compare(strings.Trim(strings.ToLower(answer), "\n"), problem.Answer) == 0 {
				score++
			}
		}

		fmt.Printf("\nYOUR SCORE: %d\n", score)
	}
}

func readCsvFile(filePath string) [][]string {
	// Load the csv file
	raw, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	// Create a new reader
	reader := csv.NewReader(raw)
	lines, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	return lines
}

func readArguments() (*string, *int) {
	csvFile := flag.String("csv", "questions.csv", "CSV file with the question for the quiz, they must be (question, value) formated.")
	timer := flag.Int("timer", 30, "Time limit for the quiz.")

	flag.Parse()

	return csvFile, timer
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problems[i] = Problem{line[0], line[1]}
	}

	return problems
}
