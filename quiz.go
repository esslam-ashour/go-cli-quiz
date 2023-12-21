package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	problemsFile    = "problems.csv"
	quizDuration    = 30
	shuffleProblems = true
	cyan            = color.New(color.FgCyan)
)

func readProblems(problemsFile string) ([][]string, error) {
	file, err := os.Open(problemsFile)
	if err != nil {
		return nil, fmt.Errorf("error while reading file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading records: %v", err)
	}

	if shuffleProblems {
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}
	return records, nil
}

func askQuestions(problems [][]string, timerExpired chan bool) (int, int) {
	correctCount := 0
	wrongCount := 0

	timer := time.NewTimer(time.Duration(quizDuration) * time.Second)

	for qid, record := range problems {
		question := record[0]
		answer := record[1]

		var userAnswer string
		fmt.Printf("Question %d: %s\n", qid+1, question)
		fmt.Println("Answer: ")

		go func() {
			fmt.Scanln(&userAnswer)
			timerExpired <- true
		}()

		userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))

		select {
		case <-timer.C:
			fmt.Println("Time's up! Quiz is over.")
			return correctCount, wrongCount
		case <-timerExpired:
			if userAnswer == answer {
				correctCount++
			} else {
				wrongCount++
			}
		}

	}
	return correctCount, wrongCount
}

func startQuiz(problemsFile string) {
	problems, err := readProblems(problemsFile)
	if err != nil {
		log.Fatal(err)
	}

	cyan.Println(">> Press any key to start the quiz <<")
	fmt.Scanln()

	timerExpired := make(chan bool)
	correct, wrong := askQuestions(problems, timerExpired)
	fmt.Printf("You got %d correct answers and %d wrong answers\n", correct, wrong)
}

func main() {
	startQuiz(problemsFile)
}
